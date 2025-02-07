package seeder

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"perizinan/model"
	"reflect"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Permit struct {
	No                    int
	Pemohon               string
	Perusahaan            string
	Alamat                string
	SumberAir             string
	Long0                 int
	Long1                 int
	Long2                 int
	Lat0                  int
	Lat1                  int
	Lat2                  int
	Desa                  string
	Kecamatan             string
	KabKota               string
	JenisUsaha            string
	DebitIzin             int
	TanggalSuratKeputusan string
	NomorSuratKeputusan   string
	MasaBerlakuIzin       string
	Status                string
	Perpanjangan          string
	Keterangan            string
}

func SeederSKPerizinan(db *gorm.DB, filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ','

	// Read the header
	_, err = reader.Read()
	if err != nil {
		log.Fatal("Error reading the header: ", err)
	}

	var permits []model.SKPerizinan

	for {
		record, err := reader.Read()
		if err != nil {
			break
		}

		latitude := convertDMStoDD(record[5], record[6], record[7], "W")
		longitude := convertDMStoDD(record[8], record[9], record[10], "N")

		debitIzin, _ := strconv.ParseFloat(record[15], 64)

		permit := model.SKPerizinan{
			ID:                model.SKPerizinanID(uuid.New().String()),
			Pemohon:           record[1],
			PerusahaanPemohon: record[2],
			AlamatPemohon:     record[3],
			SumberAir:         record[4],

			KoordinatDiDalamSK: model.NewGeometryPoint(longitude, latitude),

			Desa:      record[11],
			Kecamatan: record[12],
			KabKota:   record[13],

			JenisUsaha:             record[14],
			KuotaAirDalamSK:        debitIzin,
			TanggalSK:              record[16],
			NoSK:                   model.NomorSK(record[17]),
			MasaBerlakuSK:          record[18],
			Status:                 model.SKPerizinanStatus(record[19]),
			Perpanjangan:           record[20],
			KetentuanTeknisLainnya: record[21],
		}

		permits = append(permits, permit)
	}

	err = insertPermits(db, permits)
	if err != nil {
		return fmt.Errorf("error inserting permits: %v", err)
	}

	return nil
}

func insertPermits(db *gorm.DB, permits []model.SKPerizinan) error {
	tags := getTags[model.SKPerizinan]("no_sk")

	err := db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "no_sk"}},
		DoUpdates: clause.AssignmentColumns(tags),
	}).Create(&permits).Error

	if err != nil {
		return fmt.Errorf("error creating permits: %v", err)
	}

	return nil

}

func getTags[T any](excludes ...string) []string {
	var t T
	v := reflect.ValueOf(t)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	typ := v.Type()
	if typ.Kind() != reflect.Struct {
		return nil
	}

	// Create a map for quick lookup of excluded fields
	excludeMap := make(map[string]bool)
	for _, exclude := range excludes {
		excludeMap[exclude] = true
	}

	var tags []string
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		tag := field.Tag.Get("json")
		if tag != "" {
			// Split the tag on comma and take the first part
			// This handles cases like `json:"name,omitempty"`
			tagParts := strings.Split(tag, ",")
			tagName := tagParts[0]

			// Check if the tag should be excluded
			if !excludeMap[tagName] {
				tags = append(tags, tagName)
			}
		}
	}

	return tags
}

func convertDMStoDD(degrees, minutes, seconds string, direction string) float64 {
	deg, _ := strconv.ParseFloat(degrees, 64)
	min, _ := strconv.ParseFloat(minutes, 64)
	sec, _ := strconv.ParseFloat(seconds, 64)

	decimal := deg + (min / 60) + (sec / 3600)

	if direction == "S" || direction == "W" {
		decimal = -decimal
	}

	return decimal
}
