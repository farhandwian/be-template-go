package wiring

import (
	"dashboard/seeder"
	"shared/model"

	"gorm.io/gorm"
)

func InitSeeder(db *gorm.DB) {

	seeder.Seeder[model.Bendung](db, "infra_4_Bendung.json", 50)
	seeder.Seeder[model.Embung](db, "infra_5_Embung.json", 50)
	seeder.Seeder[model.AirTanah](db, "infra_6_Air Tanah.json", 50)
	seeder.Seeder[model.AirBaku](db, "infra_7_Air Baku.json", 50)
	seeder.Seeder[model.PengendaliSedimen](db, "infra_8_Pengendali Sedimen.json", 50)
	seeder.Seeder[model.PengamanPantai](db, "infra_12_Pengaman Pantai.json", 50)
	//seeder.SeederWithoutGeoJSON[model.CurahHujan](db, "infra_14_CURAH HUJAN.json", 50)
	//seeder.SeederWithoutGeoJSON[model.DugaAir](db, "infra_15_DUGA AIR.json", 50)
	//seeder.SeederWithoutGeoJSON[model.Klimatologi](db, "infra_16_KLIMATOLOGI.json", 50)
	seeder.Seeder[model.Sumur](db, "infra_26_Sumur.json", 50)
	seeder.Seeder[model.PahAbsah](db, "infra_25_PAH_ABSAH.json", 50)
	seeder.Seeder[model.Intake](db, "infra_27_Intake.json", 50)
	seeder.Seeder[model.Danau](db, "lake.json", 50)
	seeder.Seeder[model.Bendungan](db, "dam.json", 50)

}
