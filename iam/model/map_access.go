package model

import (
	"fmt"
	"strings"
)

type MapAccessID string

type MapAccessTypeEnum string

const (
	CREATE MapAccessTypeEnum = "CREATE"
	READ   MapAccessTypeEnum = "READ"
	UPDATE MapAccessTypeEnum = "UPDATE"
	DELETE MapAccessTypeEnum = "DELETE"
)

type MapAccess struct {
	ID          MapAccessID       `json:"id"`
	Access      Access            `json:"-"`
	Description string            `json:"description"`
	Group       string            `json:"group"`
	Type        MapAccessTypeEnum `json:"type"`
	Enabled     bool              `json:"enabled"`
	Name        string            `json:"-"`
}

func FindMapAccessByID(id MapAccessID) (*MapAccess, error) {
	mapAccesses := GetMapAccess()

	for _, access := range mapAccesses {
		if access.ID == id {
			return &access, nil
		}
	}

	return nil, fmt.Errorf("MapAccess with ID %s not found", id)
}

func FindMapAccessByName(name string) (*MapAccess, error) {

	if strings.TrimSpace(name) == "" {
		return nil, fmt.Errorf("name must not empty")
	}

	mapAccesses := GetMapAccess()

	for _, access := range mapAccesses {
		if access.Name == name {
			return &access, nil
		}
	}

	return nil, fmt.Errorf("MapAccess with Name %s not found", name)
}

func MapAccessIDsToAccess(ids []MapAccessID) []Access {
	mapAccesses := GetMapAccess()
	accessMap := make(map[MapAccessID]Access)

	// Create a map for quick lookup
	for _, access := range mapAccesses {
		accessMap[access.ID] = access.Access
	}

	// Convert IDs to Access
	result := make([]Access, 0, len(ids))
	for _, id := range ids {
		if access, ok := accessMap[id]; ok {
			result = append(result, access)
		}
	}

	return result
}

func GetMapAccess() []MapAccess {
	return accessMap
}

func GenerateMapAccess() []MapAccess {

	entries := []struct {
		Type        MapAccessTypeEnum
		Group       string
		Description string
	}{
		{READ, "Dashboard", "Data Perangkat"},
		{READ, "Dashboard", "Data Si JagaCai"},
		{READ, "Dashboard", "Data Autonomous Drone Patrolling System"},
		{READ, "Dashboard", "Data Status"},
		{READ, "Dashboard", "Tabel Data Kondisi Pemenuhan Air Irigasi"},
		{READ, "Dashboard", "Tabel Daftar Pintu Dengan Debit Tidak Terpenuhi"},
		{READ, "Dashboard", "Chart Monitoring Aktivitas"},

		{READ, "Si JagaCai", "Table Rekomtek"},
		{READ, "Si JagaCai", "Detail Rekomtek"},
		{READ, "Si JagaCai", "Sistem Pemantauan-Tabel Daftar Laporan"},

		{READ, "Si JagaCai", "Sistem Pengawasan-Tabel Daftar Kemungkinan Objek Pengguna Air Tidak Berizin"},
		{DELETE, "Si JagaCai", "Sistem Pengawasan-Tabel Daftar Kemungkinan Objek Pengguna Air Tidak Berizin"},

		{READ, "Si JagaCai", "Sistem Pemantauan-Peta"},

		{READ, "Pintu Air", "Tabel Daftar Pintu"},
		{READ, "Pintu Air", "Detail Pintu Air - Data Petugas"},
		{READ, "Pintu Air", "Detail Pintu Air - Data Status"},
		{READ, "Pintu Air", "Detail Pintu Air - CCTV"},

		{READ, "Pintu Air", "Detail Pintu Air - Pengontrolan Pintu Air"},
		{UPDATE, "Pintu Air", "Detail Pintu Air - Pengontrolan Pintu Air"},

		{READ, "Pintu Air", "Detail Pintu Air - Pengontrolan Sensor"},
		{UPDATE, "Pintu Air", "Detail Pintu Air - Pengontrolan Sensor"},

		{READ, "Pintu Air", "Detail Pintu Air - Pengontrolan Security Relay"},
		{UPDATE, "Pintu Air", "Detail Pintu Air - Pengontrolan Security Relay"},

		{READ, "Pintu Air", "Detail Pintu Air - Table data Riwayat Pengoperasian Pintu Air"},
		{READ, "Pintu Air", "Detail Pintu Air - Table data Jadwal Pengontrolan Pintu Air"},

		{CREATE, "Master Data", "Daftar Aset"},
		{READ, "Master Data", "Daftar Aset"},
		{UPDATE, "Master Data", "Daftar Aset"},
		{DELETE, "Master Data", "Daftar Aset"},

		{CREATE, "Master Data", "Daftar Proyek"},
		{READ, "Master Data", "Daftar Proyek"},
		{UPDATE, "Master Data", "Daftar Proyek"},
		{DELETE, "Master Data", "Daftar Proyek"},

		{CREATE, "Master Data", "Daftar Kepegawaian"},
		{READ, "Master Data", "Daftar Kepegawaian"},
		{UPDATE, "Master Data", "Daftar Kepegawaian"},
		{DELETE, "Master Data", "Daftar Kepegawaian"},

		{CREATE, "Master Data", "Daftar JDIH"},
		{READ, "Master Data", "Daftar JDIH"},
		{UPDATE, "Master Data", "Daftar JDIH"},
		{DELETE, "Master Data", "Daftar JDIH"},

		{CREATE, "Konfigurasi Pengetahuan AI", "Kamus"},
		{READ, "Konfigurasi Pengetahuan AI", "Kamus"},
		{UPDATE, "Konfigurasi Pengetahuan AI", "Kamus"},
		{DELETE, "Konfigurasi Pengetahuan AI", "Kamus"},

		{CREATE, "Konfigurasi Pengetahuan AI", "Dokumen"},
		{READ, "Konfigurasi Pengetahuan AI", "Dokumen"},
		{UPDATE, "Konfigurasi Pengetahuan AI", "Dokumen"},
		{DELETE, "Konfigurasi Pengetahuan AI", "Dokumen"},

		{CREATE, "Sistem Peringatan", "Konfigurasi"},
		{READ, "Sistem Peringatan", "Konfigurasi"},
		{UPDATE, "Sistem Peringatan", "Konfigurasi"},
		{DELETE, "Sistem Peringatan", "Konfigurasi"},

		{CREATE, "Manajemen Pengguna", "Daftar Pengguna"},
		{READ, "Manajemen Pengguna", "Daftar Pengguna"},
		{UPDATE, "Manajemen Pengguna", "Daftar Pengguna"},
		{DELETE, "Manajemen Pengguna", "Daftar Pengguna"},

		{UPDATE, "Manajemen Pengguna", "Hak Akses"},
		{UPDATE, "Manajemen Pengguna", "Reset Kata Sandi"},
		{UPDATE, "Manajemen Pengguna", "Aktivasi Button"},
	}

	mapAccesses := []MapAccess{}
	base := 1
	digits := 0
	counter := 1

	for _, t := range entries {

		access := fmt.Sprintf("%d", base)
		if digits > 0 {
			access += strings.Repeat("0", digits)
		}

		mapAccess := MapAccess{
			ID:          MapAccessID(fmt.Sprintf("%d", counter)),
			Access:      Access(access),
			Description: t.Description,
			Group:       t.Group,
			Type:        t.Type,
		}
		mapAccesses = append(mapAccesses, mapAccess)
		counter++

		base *= 2
		if base > 8 {
			base = 1
			digits++
		}

	}

	return mapAccesses

}
