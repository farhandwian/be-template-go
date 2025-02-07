package cronjob

import (
	"context"
	"fmt"
	"log"
	"shared/middleware"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// GormCronJobState is the database model for CronJobState
type GormCronJobState struct {
	FuncID        string     `gorm:"column:func_id;primaryKey"`
	TargetDate    time.Time  `gorm:"column:target_date;not null"`
	CreatedAt     time.Time  `gorm:"column:created_at;not null"`
	FunctionType  string     `gorm:"column:function_type;not null"`
	Status        string     `gorm:"column:status;not null"`
	ErrorMessage  string     `gorm:"column:error_message"`
	Data          []byte     `gorm:"column:data;type:blob"`
	MissedDate    *time.Time `gorm:"column:missed_date"`
	CancelledDate *time.Time `gorm:"column:cancelled_date"`
}

// TableName specifies the table name for the model
func (GormCronJobState) TableName() string {
	return "cron_jobs"
}

type MariaDBStorage struct {
	db *gorm.DB
}

func NewMariaDBStorageDefault() *MariaDBStorage {

	dbHost := "localhost"
	dbPort := "3306"
	dbUser := "root"
	dbPassword := "12345"
	dbName := "test_cronjob"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	// Initialize DB connection
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	mdb, err := NewMariaDBStorage(db)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	return mdb
}

func NewMariaDBStorage(db *gorm.DB) (*MariaDBStorage, error) {
	// Auto-migrate the schema
	if err := db.AutoMigrate(&GormCronJobState{}); err != nil {
		return nil, fmt.Errorf("failed to migrate database: %v", err)
	}

	return &MariaDBStorage{db: db}, nil
}

func (s *MariaDBStorage) SaveJob(ctx context.Context, state CronJobState) error {

	gormState := GormCronJobState{
		FuncID:        string(state.FuncID),
		TargetDate:    state.TargetDate,
		CreatedAt:     state.CreatedAt,
		FunctionType:  string(state.FunctionType),
		Status:        string(state.Status),
		ErrorMessage:  state.ErrorMessage,
		Data:          state.Data,
		MissedDate:    state.MissedDate,
		CancelledDate: state.CancelledDate,
	}

	result := middleware.GetDBFromContext(ctx, s.db).Create(&gormState)
	return result.Error
}

func (s *MariaDBStorage) DeleteJob(ctx context.Context, funcID FuncID) error {
	result := middleware.GetDBFromContext(ctx, s.db).Delete(&GormCronJobState{}, "func_id = ?", string(funcID))
	return result.Error
}

func (s *MariaDBStorage) GetAllJobs(ctx context.Context) ([]CronJobState, error) {
	var gormStates []GormCronJobState
	if err := middleware.GetDBFromContext(ctx, s.db).Find(&gormStates).Error; err != nil {
		return nil, err
	}

	states := make([]CronJobState, len(gormStates))
	for i, gormState := range gormStates {
		states[i] = CronJobState{
			FuncID:        FuncID(gormState.FuncID),
			TargetDate:    gormState.TargetDate,
			CreatedAt:     gormState.CreatedAt,
			FunctionType:  FuncType(gormState.FunctionType),
			Status:        CronJobStatus(gormState.Status),
			ErrorMessage:  gormState.ErrorMessage,
			Data:          gormState.Data,
			MissedDate:    gormState.MissedDate,
			CancelledDate: gormState.CancelledDate,
		}
	}

	return states, nil
}

func (s *MariaDBStorage) GetJob(ctx context.Context, funcID FuncID) (*CronJobState, error) {
	var gormState GormCronJobState
	if err := middleware.GetDBFromContext(ctx, s.db).First(&gormState, "func_id = ?", string(funcID)).Error; err != nil {
		return nil, err
	}

	return &CronJobState{
		FuncID:        FuncID(gormState.FuncID),
		TargetDate:    gormState.TargetDate,
		CreatedAt:     gormState.CreatedAt,
		FunctionType:  FuncType(gormState.FunctionType),
		Status:        CronJobStatus(gormState.Status),
		ErrorMessage:  gormState.ErrorMessage,
		Data:          gormState.Data,
		MissedDate:    gormState.MissedDate,
		CancelledDate: gormState.CancelledDate,
	}, nil
}

func (s *MariaDBStorage) UpdateJobStatus(ctx context.Context, funcID FuncID, status CronJobStatus, errorMsg string, statusDate *time.Time) error {
	updates := map[string]interface{}{
		"status":        string(status),
		"error_message": errorMsg,
	}

	switch status {
	case Missed:
		updates["missed_date"] = statusDate
	case Cancelled:
		updates["cancelled_date"] = statusDate
	}

	result := middleware.GetDBFromContext(ctx, s.db).Model(&GormCronJobState{}).
		Where("func_id = ?", string(funcID)).
		Updates(updates)

	return result.Error
}
