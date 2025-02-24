package config

import (
	"example/model"
	"fmt"
	iamModel "iam/model"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitMariaDatabase() *gorm.DB {

	// Retrieve database connection details from environment variables
	dbUser := os.Getenv("DB_EXAMPLE_USER")
	dbPassword := os.Getenv("DB_EXAMPLE_PASSWORD")
	dbHost := os.Getenv("DB_EXAMPLE_HOST")
	dbPort := os.Getenv("DB_EXAMPLE_PORT")
	dbName := os.Getenv("DB_EXAMPLE_NAME")

	// Construct the DSN string using environment variables
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             0,           // Log all queries
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,        // Enable color
		},
	)

	_ = newLogger

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// Logger: newLogger,
	})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(
		&model.Example{},
		&iamModel.User{},
	)
	// &model.SKPerizinan{},
	// &model.LaporanPerizinan{},
	// 	&model.Employee{},
	// 	&model.Asset{},
	// 	&model.Project{},
	// 	&model.JDIH{},
	// 	&sharedModel.Danau{},
	// 	&sharedModel.Bendungan{},
	// 	&sharedModel.Bendung{},
	// 	&sharedModel.Embung{},
	// 	&sharedModel.AirTanah{},
	// 	&sharedModel.PengendaliSedimen{},
	// 	&sharedModel.PengamanPantai{},
	// 	&sharedModel.Sumur{},
	// 	&model.CurahHujan{},
	// &model.Klimatologi{},
	// 	&sharedModel.AirBaku{},
	// 	&model.DugaAir{},
	// 	&sharedModel.PahAbsah{},
	// 	&sharedModel.Intake{},
	// &model.AlarmConfig{},
	// 	&model.DoorControl{},
	// 	&model.DoorControlHistory{},
	// 	&sharedModel.ActivityMonitor{},
	// 	&iamModel.User{},
	// &model.DoorControl{},
	// &model.DoorControlHistory{},
	// &model.AlarmHistory{},
	// &model.CctvImageProcessing{},
	// )

	// Verify the connection
	sqlDB, err := db.DB()
	if err != nil {
		panic("failed to get database")
	}

	err = sqlDB.Ping()
	if err != nil {
		panic("failed to ping database")
	}

	// err = createMariaDBIndex(db)
	// if err != nil {
	// 	log.Printf("failed to create indexes: %v", err)
	// }

	return db
}

func createMariaDBIndex(db *gorm.DB) error {
	err := db.Exec(`
        CREATE INDEX IF NOT EXISTS idx_waterchanneldoor_name
        ON water_channel_doors(name);
    `).Error
	if err != nil {
		return fmt.Errorf("error creating index on water_channel_doors: %w", err)
	}

	err = db.Exec(`CREATE INDEX IF NOT EXISTS idx_water_channel_device_category
	ON water_channel_devices(category);`).Error
	if err != nil {
		return fmt.Errorf("error creating index on water_channel_devices: %w", err)
	}

	err = db.Exec(`
	CREATE INDEX IF NOT EXISTS idx_water_channel_device_water_channel_door_id
		ON water_channel_devices(water_channel_door_id);
   	`).Error
	if err != nil {
		return fmt.Errorf("error creating index on water_channel_devices: %w", err)
	}

	err = db.Exec(`
	CREATE INDEX IF NOT EXISTS idx_water_channel_officer_water_channel_door_id 
	ON water_channel_officers(water_channel_door_id);
  	`).Error
	if err != nil {
		return fmt.Errorf("error creating index on water_channel_officers: %w", err)
	}
	return nil
}
