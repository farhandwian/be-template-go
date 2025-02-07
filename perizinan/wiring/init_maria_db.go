package wiring

import (
	"fmt"
	"log"
	"os"
	"perizinan/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitMariaDatabase() *gorm.DB {

	// Retrieve database connection details from environment variables
	dbUser := os.Getenv("DB_PERIZINAN_USER")
	dbPassword := os.Getenv("DB_PERIZINAN_PASSWORD")
	dbHost := os.Getenv("DB_PERIZINAN_HOST")
	dbPort := os.Getenv("DB_PERIZINAN_PORT")
	dbName := os.Getenv("DB_PERIZINAN_NAME")

	// Construct the DSN string using environment variables
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	// fmt.Printf("Connection String : %s\n", dsn)

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

	err = db.AutoMigrate(
		&model.SKPerizinan{},
		&model.LaporanPerizinan{},
	)
	if err != nil {
		panic("failed to migrate database: " + err.Error())
	}

	// seeder
	// seeder.SeederSKPerizinan(db, "./seeder/seeder_sk.csv")

	// seeder.Seeder[model.Bendung](db, "infra_4_Bendung.json", 50)
	// seeder.Seeder[model.Embung](db, "infra_5_Embung.json", 50)
	// seeder.Seeder[model.AirTanah](db, "infra_6_Air Tanah.json", 50)
	// seeder.Seeder[model.AirBaku](db, "infra_7_Air Baku.json", 50)
	// seeder.Seeder[model.PengendaliSedimen](db, "infra_8_Pengendali Sedimen.json", 50)
	// seeder.Seeder[model.PengamanPantai](db, "infra_12_Pengaman Pantai.json", 50)
	// seeder.SeederWithoutGeoJSON[model.CurahHujan](db, "infra_14_CURAH HUJAN.json", 50)
	// seeder.SeederWithoutGeoJSON[model.DugaAir](db, "infra_15_DUGA AIR.json", 50)
	// seeder.SeederWithoutGeoJSON[model.Klimatologi](db, "infra_16_KLIMATOLOGI.json", 50)
	// seeder.Seeder[model.Sumur](db, "infra_26_Sumur.json", 50)
	// seeder.Seeder[model.PahAbsah](db, "infra_25_PAH_ABSAH.json", 50)
	// seeder.Seeder[model.Intake](db, "infra_27_Intake.json", 50)
	// seeder.Seeder[model.Danau](db, "lake.json", 50)
	// seeder.Seeder[model.Bendungan](db, "dam.json", 50)

	// Verify the connection
	sqlDB, err := db.DB()
	if err != nil {
		panic("failed to get database: " + err.Error())
	}

	err = sqlDB.Ping()
	if err != nil {
		panic("failed to ping database: " + err.Error())
	}

	return db
}
