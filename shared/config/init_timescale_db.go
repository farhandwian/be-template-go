package config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitTimeSeriesDatabase() *gorm.DB {
	var err error

	tsdbUser := os.Getenv("TSDB_USER")
	tsdbPassword := os.Getenv("TSDB_PASSWORD")
	tsdbHost := os.Getenv("TSDB_HOST")
	tsdbPort := os.Getenv("TSDB_PORT")
	tsdbName := os.Getenv("TSDB_NAME")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		tsdbHost,
		tsdbUser,
		tsdbPassword,
		tsdbName,
		tsdbPort,
	)

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

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect timescale database")
	}

	// Set connection pool settings
	tsDB, err := db.DB()
	if err != nil {
		panic("failed to get database")
	}

	err = tsDB.Ping()
	if err != nil {
		panic("failed to ping database")
	}

	err = createIndexes(db)
	if err != nil {
		log.Printf("failed to create indexes: %v", err)
	}

	err = InitializeMaterializedViews(db)
	if err != nil {
		log.Printf("failed to create materialized views: %v", err)
	}
	return db
}

// Create necessary indexes for performance optimization
func createIndexes(db *gorm.DB) error {
	//err := db.Exec(`
	//	CREATE INDEX IF NOT EXISTS idx_water_gates_channel_device_timestamp_gate_level
	//ON water_gates (water_channel_door_id, device_id, timestamp DESC, gate_level);
	//`).Error
	//if err != nil {
	//	return fmt.Errorf("failed to create composite index: %w", err)
	//}

	err := db.Exec(`
		CREATE INDEX IF NOT EXISTS idx_water_surface_elevations_channel_timestamp_desc 
		ON water_surface_elevations (water_channel_door_id, timestamp DESC)
	`).Error
	if err != nil {
		return fmt.Errorf("failed to create optimized index: %w", err)
	}

	err = db.Exec(`
		CREATE INDEX IF NOT EXISTS idx_actual_debits_channel_timestamp 
		ON actual_debits (water_channel_door_id, timestamp DESC)
	`).Error
	if err != nil {
		return fmt.Errorf("failed to create composite index: %w", err)
	}

	return nil
}

func InitializeMaterializedViews(db *gorm.DB) error {
	// 1. Check and create materialized view for `latest_actual_debits`
	var existsDebit bool
	err := db.Raw(`
        SELECT EXISTS (
			SELECT 1
			FROM pg_matviews
			WHERE matviewname = 'latest_actual_debits'
		);
    `).Scan(&existsDebit).Error

	if err != nil {
		return fmt.Errorf("error checking debit materialized view: %v", err)
	}

	if !existsDebit {
		// Create materialized view for debit
		err = db.Exec(`
            CREATE MATERIALIZED VIEW latest_actual_debits AS
			SELECT DISTINCT ON (water_channel_door_id)
				timestamp,
				water_channel_door_id,
				ROUND(actual_debit) as latest_debit
			FROM actual_debit_data
			WHERE timestamp > NOW() - INTERVAL '1 hour'
			ORDER BY water_channel_door_id, timestamp DESC;
        `).Error

		if err != nil {
			return fmt.Errorf("error creating debit materialized view: %v", err)
		}

		log.Println("Successfully created debit materialized view and policy")
	} else {
		log.Println("Debit materialized view already exists, skipping")
	}

	// 2. Check and create materialized view for `latest_water_levels`
	var existsWater bool
	err = db.Raw(`
        SELECT EXISTS (
			SELECT 1
			FROM pg_matviews
			WHERE matviewname = 'latest_water_levels'
		);
    `).Scan(&existsWater).Error

	if err != nil {
		return fmt.Errorf("error checking water levels materialized view: %v", err)
	}

	if !existsWater {
		// Create materialized view for water levels
		err = db.Exec(`
            CREATE MATERIALIZED VIEW latest_water_levels AS
			SELECT DISTINCT ON (water_channel_door_id)
				timestamp,
				water_channel_door_id,
				ROUND(water_level,1) as latest_level,
				status as latest_status
			FROM water_surface_elevation_data
			WHERE timestamp > NOW() - INTERVAL '1 days'
			ORDER BY water_channel_door_id, timestamp DESC;
        `).Error

		if err != nil {
			return fmt.Errorf("error creating water levels materialized view: %v", err)
		}

		log.Println("Successfully created water levels materialized view")
	} else {
		log.Println("Water levels materialized view already exists, skipping")
	}

	return nil
}
