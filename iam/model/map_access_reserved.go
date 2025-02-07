package model

const (
	ANONYMOUS         Access = Access("0")
	ADMIN_OPERATION   Access = Access("1")
	DEFAULT_OPERATION Access = Access("0")

	SERVICE_OPERATION Access = Access("4")
	WEBHOOK_OPERATION Access = Access("8")

	// ASSET_READ    Access = Access("10")
	// ASSET_WRITE   Access = Access("20")
	// PROJECT_READ  Access = Access("40")
	// PROJECT_WRITE Access = Access("80")

	// EMPLOYEE_READ  Access = Access("100")
	// EMPLOYEE_WRITE Access = Access("200")
	// JDIH_READ      Access = Access("400")
	// JDIH_WRITE     Access = Access("800")

	// ALARM_READ         Access = Access("1000")
	// ALARM_WRITE        Access = Access("2000")
	// DOOR_CONTROL_READ  Access = Access("4000")
	// DOOR_CONTROL_WRITE Access = Access("8000")
)

var accessMap = []MapAccess{
	{ID: "1", Access: DASHBOARD_DATA_PERANGKAT_READ, Description: "Data Perangkat", Group: "Dashboard", Type: READ},
	{ID: "2", Access: DASHBOARD_DATA_SI_JAGACAI_READ, Description: "Data Si JagaCai", Group: "Dashboard", Type: READ},
	{ID: "3", Access: DASHBOARD_DATA_AUTONOMOUS_DRONE_PATROLLING_SYSTEM_READ, Description: "Data Autonomous Drone Patrolling System", Group: "Dashboard", Type: READ},
	{ID: "4", Access: DASHBOARD_DATA_STATUS_READ, Description: "Data Status", Group: "Dashboard", Type: READ},
	{ID: "5", Access: DASHBOARD_TABEL_DATA_KONDISI_PEMENUHAN_AIR_IRIGASI_READ, Description: "Tabel Data Kondisi Pemenuhan Air Irigasi", Group: "Dashboard", Type: READ},
	{ID: "6", Access: DASHBOARD_TABEL_DAFTAR_PINTU_DENGAN_DEBIT_TIDAK_TERPENUHI_READ, Description: "Tabel Daftar Pintu Dengan Debit Tidak Terpenuhi", Group: "Dashboard", Type: READ},
	{ID: "7", Access: DASHBOARD_MONITORING_AKTIVITAS_READ, Description: "Monitoring Aktivitas", Group: "Dashboard", Type: READ},
	{ID: "8", Access: SI_JAGACAI_TABLE_REKOMTEK_READ, Description: "Table Rekomtek", Group: "Si JagaCai", Type: READ, Name: "sk_read"},
	{ID: "9", Access: SI_JAGACAI_DETAIL_REKOMTEK_READ, Description: "Detail Rekomtek", Group: "Si JagaCai", Type: READ},
	{ID: "10", Access: SI_JAGACAI_SISTEM_PEMANTAUAN_TABEL_DAFTAR_LAPORAN_READ, Description: "Sistem Pemantauan-Tabel Daftar Laporan", Group: "Si JagaCai", Type: READ, Name: "perizinan_read"},
	{ID: "11", Access: SI_JAGACAI_SISTEM_PENGAWASAN_TABEL_DAFTAR_KEMUNGKINAN_OBJEK_PENGGUNA_AIR_TIDAK_BERIZIN_READ, Description: "Sistem Pengawasan-Tabel Daftar Kemungkinan Objek Pengguna Air Tidak Berizin", Group: "Si JagaCai", Type: READ},
	{ID: "12", Access: SI_JAGACAI_SISTEM_PENGAWASAN_TABEL_DAFTAR_KEMUNGKINAN_OBJEK_PENGGUNA_AIR_TIDAK_BERIZIN_DELETE, Description: "Sistem Pengawasan-Tabel Daftar Kemungkinan Objek Pengguna Air Tidak Berizin", Group: "Si JagaCai", Type: DELETE},
	{ID: "13", Access: SI_JAGACAI_SISTEM_PEMANTAUAN_PETA_READ, Description: "Sistem Pemantauan-Peta", Group: "Si JagaCai", Type: READ},
	{ID: "14", Access: PINTU_AIR_TABEL_DAFTAR_PINTU_READ, Description: "Tabel Daftar Pintu", Group: "Pintu Air", Type: READ},
	{ID: "15", Access: PINTU_AIR_DETAIL_PINTU_AIR_DATA_PETUGAS_READ, Description: "Detail Pintu Air - Data Petugas", Group: "Pintu Air", Type: READ},
	{ID: "16", Access: PINTU_AIR_DETAIL_PINTU_AIR_DATA_STATUS_READ, Description: "Detail Pintu Air - Data Status", Group: "Pintu Air", Type: READ},
	{ID: "17", Access: PINTU_AIR_DETAIL_PINTU_AIR_CCTV_READ, Description: "Detail Pintu Air - CCTV", Group: "Pintu Air", Type: READ},
	{ID: "18", Access: PINTU_AIR_DETAIL_PINTU_AIR_PENGONTROLAN_PINTU_AIR_READ, Description: "Detail Pintu Air - Pengontrolan Pintu Air", Group: "Pintu Air", Type: READ},
	{ID: "19", Access: PINTU_AIR_DETAIL_PINTU_AIR_PENGONTROLAN_PINTU_AIR_UPDATE, Description: "Detail Pintu Air - Pengontrolan Pintu Air", Group: "Pintu Air", Type: UPDATE},
	{ID: "20", Access: PINTU_AIR_DETAIL_PINTU_AIR_PENGONTROLAN_SENSOR_READ, Description: "Detail Pintu Air - Pengontrolan Sensor", Group: "Pintu Air", Type: READ},
	{ID: "21", Access: PINTU_AIR_DETAIL_PINTU_AIR_PENGONTROLAN_SENSOR_UPDATE, Description: "Detail Pintu Air - Pengontrolan Sensor", Group: "Pintu Air", Type: UPDATE},
	{ID: "22", Access: PINTU_AIR_DETAIL_PINTU_AIR_PENGONTROLAN_SECURITY_RELAY_READ, Description: "Detail Pintu Air - Pengontrolan Security Relay", Group: "Pintu Air", Type: READ},
	{ID: "23", Access: PINTU_AIR_DETAIL_PINTU_AIR_PENGONTROLAN_SECURITY_RELAY_UPDATE, Description: "Detail Pintu Air - Pengontrolan Security Relay", Group: "Pintu Air", Type: UPDATE},
	{ID: "24", Access: PINTU_AIR_DETAIL_PINTU_AIR_TABLE_DATA_RIWAYAT_PENGOPERASIAN_PINTU_AIR_READ, Description: "Detail Pintu Air - Table data Riwayat Pengoperasian Pintu Air", Group: "Pintu Air", Type: READ},
	{ID: "25", Access: PINTU_AIR_DETAIL_PINTU_AIR_TABLE_DATA_JADWAL_PENGONTROLAN_PINTU_AIR_READ, Description: "Detail Pintu Air - Table data Jadwal Pengontrolan Pintu Air", Group: "Pintu Air", Type: READ},
	{ID: "26", Access: MASTER_DATA_DAFTAR_ASET_CREATE, Description: "Daftar Aset", Group: "Master Data", Type: CREATE},
	{ID: "27", Access: MASTER_DATA_DAFTAR_ASET_READ, Description: "Daftar Aset", Group: "Master Data", Type: READ},
	{ID: "28", Access: MASTER_DATA_DAFTAR_ASET_UPDATE, Description: "Daftar Aset", Group: "Master Data", Type: UPDATE},
	{ID: "29", Access: MASTER_DATA_DAFTAR_ASET_DELETE, Description: "Daftar Aset", Group: "Master Data", Type: DELETE},
	{ID: "30", Access: MASTER_DATA_DAFTAR_PROYEK_CREATE, Description: "Daftar Proyek", Group: "Master Data", Type: CREATE},
	{ID: "31", Access: MASTER_DATA_DAFTAR_PROYEK_READ, Description: "Daftar Proyek", Group: "Master Data", Type: READ},
	{ID: "32", Access: MASTER_DATA_DAFTAR_PROYEK_UPDATE, Description: "Daftar Proyek", Group: "Master Data", Type: UPDATE},
	{ID: "33", Access: MASTER_DATA_DAFTAR_PROYEK_DELETE, Description: "Daftar Proyek", Group: "Master Data", Type: DELETE},
	{ID: "34", Access: MASTER_DATA_DAFTAR_KEPEGAWAIAN_CREATE, Description: "Daftar Kepegawaian", Group: "Master Data", Type: CREATE},
	{ID: "35", Access: MASTER_DATA_DAFTAR_KEPEGAWAIAN_READ, Description: "Daftar Kepegawaian", Group: "Master Data", Type: READ},
	{ID: "36", Access: MASTER_DATA_DAFTAR_KEPEGAWAIAN_UPDATE, Description: "Daftar Kepegawaian", Group: "Master Data", Type: UPDATE},
	{ID: "37", Access: MASTER_DATA_DAFTAR_KEPEGAWAIAN_DELETE, Description: "Daftar Kepegawaian", Group: "Master Data", Type: DELETE},
	{ID: "38", Access: MASTER_DATA_DAFTAR_JDIH_CREATE, Description: "Daftar JDIH", Group: "Master Data", Type: CREATE},
	{ID: "39", Access: MASTER_DATA_DAFTAR_JDIH_READ, Description: "Daftar JDIH", Group: "Master Data", Type: READ},
	{ID: "40", Access: MASTER_DATA_DAFTAR_JDIH_UPDATE, Description: "Daftar JDIH", Group: "Master Data", Type: UPDATE},
	{ID: "41", Access: MASTER_DATA_DAFTAR_JDIH_DELETE, Description: "Daftar JDIH", Group: "Master Data", Type: DELETE},
	{ID: "42", Access: KONFIGURASI_PENGETAHUAN_AI_KAMUS_CREATE, Description: "Kamus", Group: "Konfigurasi Pengetahuan AI", Type: CREATE, Name: "dict_ai_create"},
	{ID: "43", Access: KONFIGURASI_PENGETAHUAN_AI_KAMUS_READ, Description: "Kamus", Group: "Konfigurasi Pengetahuan AI", Type: READ, Name: "dict_ai_read"},
	{ID: "44", Access: KONFIGURASI_PENGETAHUAN_AI_KAMUS_UPDATE, Description: "Kamus", Group: "Konfigurasi Pengetahuan AI", Type: UPDATE, Name: "dict_ai_update"},
	{ID: "45", Access: KONFIGURASI_PENGETAHUAN_AI_KAMUS_DELETE, Description: "Kamus", Group: "Konfigurasi Pengetahuan AI", Type: DELETE, Name: "dict_ai_delete"},
	{ID: "46", Access: KONFIGURASI_PENGETAHUAN_AI_DOKUMEN_CREATE, Description: "Dokumen", Group: "Konfigurasi Pengetahuan AI", Type: CREATE, Name: "doc_ai_create"},
	{ID: "47", Access: KONFIGURASI_PENGETAHUAN_AI_DOKUMEN_READ, Description: "Dokumen", Group: "Konfigurasi Pengetahuan AI", Type: READ, Name: "doc_ai_read"},
	{ID: "48", Access: KONFIGURASI_PENGETAHUAN_AI_DOKUMEN_UPDATE, Description: "Dokumen", Group: "Konfigurasi Pengetahuan AI", Type: UPDATE, Name: "doc_ai_update"},
	{ID: "49", Access: KONFIGURASI_PENGETAHUAN_AI_DOKUMEN_DELETE, Description: "Dokumen", Group: "Konfigurasi Pengetahuan AI", Type: DELETE, Name: "doc_ai_delete"},
	{ID: "50", Access: SISTEM_PERINGATAN_KONFIGURASI_CREATE, Description: "Konfigurasi", Group: "Sistem Peringatan", Type: CREATE},
	{ID: "51", Access: SISTEM_PERINGATAN_KONFIGURASI_READ, Description: "Konfigurasi", Group: "Sistem Peringatan", Type: READ},
	{ID: "52", Access: SISTEM_PERINGATAN_KONFIGURASI_UPDATE, Description: "Konfigurasi", Group: "Sistem Peringatan", Type: UPDATE},
	{ID: "53", Access: SISTEM_PERINGATAN_KONFIGURASI_DELETE, Description: "Konfigurasi", Group: "Sistem Peringatan", Type: DELETE},
	{ID: "54", Access: MANAJEMEN_PENGGUNA_DAFTAR_PENGGUNA_CREATE, Description: "Daftar Pengguna", Group: "Manajemen Pengguna", Type: CREATE},
	{ID: "55", Access: MANAJEMEN_PENGGUNA_DAFTAR_PENGGUNA_READ, Description: "Daftar Pengguna", Group: "Manajemen Pengguna", Type: READ},
	{ID: "56", Access: MANAJEMEN_PENGGUNA_DAFTAR_PENGGUNA_UPDATE, Description: "Daftar Pengguna", Group: "Manajemen Pengguna", Type: UPDATE},
	{ID: "57", Access: MANAJEMEN_PENGGUNA_DAFTAR_PENGGUNA_DELETE, Description: "Daftar Pengguna", Group: "Manajemen Pengguna", Type: DELETE},
	{ID: "58", Access: MANAJEMEN_PENGGUNA_HAK_AKSES_UPDATE, Description: "Hak Akses", Group: "Manajemen Pengguna", Type: UPDATE},
	{ID: "59", Access: MANAJEMEN_PENGGUNA_RESET_KATA_SANDI_UPDATE, Description: "Reset Kata Sandi", Group: "Manajemen Pengguna", Type: UPDATE},
	{ID: "60", Access: MANAJEMEN_PENGGUNA_AKTIVASI_BUTTON_UPDATE, Description: "Aktivasi Button", Group: "Manajemen Pengguna", Type: UPDATE},
}

const (
	DASHBOARD_DATA_PERANGKAT_READ                                                                 Access = "1"
	DASHBOARD_DATA_SI_JAGACAI_READ                                                                Access = "2"
	DASHBOARD_DATA_AUTONOMOUS_DRONE_PATROLLING_SYSTEM_READ                                        Access = "4"
	DASHBOARD_DATA_STATUS_READ                                                                    Access = "8"
	DASHBOARD_TABEL_DATA_KONDISI_PEMENUHAN_AIR_IRIGASI_READ                                       Access = "10"
	DASHBOARD_TABEL_DAFTAR_PINTU_DENGAN_DEBIT_TIDAK_TERPENUHI_READ                                Access = "20"
	DASHBOARD_MONITORING_AKTIVITAS_READ                                                           Access = "40"
	SI_JAGACAI_TABLE_REKOMTEK_READ                                                                Access = "80"
	SI_JAGACAI_DETAIL_REKOMTEK_READ                                                               Access = "100"
	SI_JAGACAI_SISTEM_PEMANTAUAN_TABEL_DAFTAR_LAPORAN_READ                                        Access = "200"
	SI_JAGACAI_SISTEM_PENGAWASAN_TABEL_DAFTAR_KEMUNGKINAN_OBJEK_PENGGUNA_AIR_TIDAK_BERIZIN_READ   Access = "400"
	SI_JAGACAI_SISTEM_PENGAWASAN_TABEL_DAFTAR_KEMUNGKINAN_OBJEK_PENGGUNA_AIR_TIDAK_BERIZIN_DELETE Access = "800"
	SI_JAGACAI_SISTEM_PEMANTAUAN_PETA_READ                                                        Access = "1000"
	PINTU_AIR_TABEL_DAFTAR_PINTU_READ                                                             Access = "2000"
	PINTU_AIR_DETAIL_PINTU_AIR_DATA_PETUGAS_READ                                                  Access = "4000"
	PINTU_AIR_DETAIL_PINTU_AIR_DATA_STATUS_READ                                                   Access = "8000"
	PINTU_AIR_DETAIL_PINTU_AIR_CCTV_READ                                                          Access = "10000"
	PINTU_AIR_DETAIL_PINTU_AIR_PENGONTROLAN_PINTU_AIR_READ                                        Access = "20000"
	PINTU_AIR_DETAIL_PINTU_AIR_PENGONTROLAN_PINTU_AIR_UPDATE                                      Access = "40000"
	PINTU_AIR_DETAIL_PINTU_AIR_PENGONTROLAN_SENSOR_READ                                           Access = "80000"
	PINTU_AIR_DETAIL_PINTU_AIR_PENGONTROLAN_SENSOR_UPDATE                                         Access = "100000"
	PINTU_AIR_DETAIL_PINTU_AIR_PENGONTROLAN_SECURITY_RELAY_READ                                   Access = "200000"
	PINTU_AIR_DETAIL_PINTU_AIR_PENGONTROLAN_SECURITY_RELAY_UPDATE                                 Access = "400000"
	PINTU_AIR_DETAIL_PINTU_AIR_TABLE_DATA_RIWAYAT_PENGOPERASIAN_PINTU_AIR_READ                    Access = "800000"
	PINTU_AIR_DETAIL_PINTU_AIR_TABLE_DATA_JADWAL_PENGONTROLAN_PINTU_AIR_READ                      Access = "1000000"
	MASTER_DATA_DAFTAR_ASET_CREATE                                                                Access = "2000000"
	MASTER_DATA_DAFTAR_ASET_READ                                                                  Access = "4000000"
	MASTER_DATA_DAFTAR_ASET_UPDATE                                                                Access = "8000000"
	MASTER_DATA_DAFTAR_ASET_DELETE                                                                Access = "10000000"
	MASTER_DATA_DAFTAR_PROYEK_CREATE                                                              Access = "20000000"
	MASTER_DATA_DAFTAR_PROYEK_READ                                                                Access = "40000000"
	MASTER_DATA_DAFTAR_PROYEK_UPDATE                                                              Access = "80000000"
	MASTER_DATA_DAFTAR_PROYEK_DELETE                                                              Access = "100000000"
	MASTER_DATA_DAFTAR_KEPEGAWAIAN_CREATE                                                         Access = "200000000"
	MASTER_DATA_DAFTAR_KEPEGAWAIAN_READ                                                           Access = "400000000"
	MASTER_DATA_DAFTAR_KEPEGAWAIAN_UPDATE                                                         Access = "800000000"
	MASTER_DATA_DAFTAR_KEPEGAWAIAN_DELETE                                                         Access = "1000000000"
	MASTER_DATA_DAFTAR_JDIH_CREATE                                                                Access = "2000000000"
	MASTER_DATA_DAFTAR_JDIH_READ                                                                  Access = "4000000000"
	MASTER_DATA_DAFTAR_JDIH_UPDATE                                                                Access = "8000000000"
	MASTER_DATA_DAFTAR_JDIH_DELETE                                                                Access = "10000000000"
	KONFIGURASI_PENGETAHUAN_AI_KAMUS_CREATE                                                       Access = "20000000000"
	KONFIGURASI_PENGETAHUAN_AI_KAMUS_READ                                                         Access = "40000000000"
	KONFIGURASI_PENGETAHUAN_AI_KAMUS_UPDATE                                                       Access = "80000000000"
	KONFIGURASI_PENGETAHUAN_AI_KAMUS_DELETE                                                       Access = "100000000000"
	KONFIGURASI_PENGETAHUAN_AI_DOKUMEN_CREATE                                                     Access = "200000000000"
	KONFIGURASI_PENGETAHUAN_AI_DOKUMEN_READ                                                       Access = "400000000000"
	KONFIGURASI_PENGETAHUAN_AI_DOKUMEN_UPDATE                                                     Access = "800000000000"
	KONFIGURASI_PENGETAHUAN_AI_DOKUMEN_DELETE                                                     Access = "1000000000000"
	SISTEM_PERINGATAN_KONFIGURASI_CREATE                                                          Access = "2000000000000"
	SISTEM_PERINGATAN_KONFIGURASI_READ                                                            Access = "4000000000000"
	SISTEM_PERINGATAN_KONFIGURASI_UPDATE                                                          Access = "8000000000000"
	SISTEM_PERINGATAN_KONFIGURASI_DELETE                                                          Access = "10000000000000"
	MANAJEMEN_PENGGUNA_DAFTAR_PENGGUNA_CREATE                                                     Access = "20000000000000"
	MANAJEMEN_PENGGUNA_DAFTAR_PENGGUNA_READ                                                       Access = "40000000000000"
	MANAJEMEN_PENGGUNA_DAFTAR_PENGGUNA_UPDATE                                                     Access = "80000000000000"
	MANAJEMEN_PENGGUNA_DAFTAR_PENGGUNA_DELETE                                                     Access = "100000000000000"
	MANAJEMEN_PENGGUNA_HAK_AKSES_UPDATE                                                           Access = "200000000000000"
	MANAJEMEN_PENGGUNA_RESET_KATA_SANDI_UPDATE                                                    Access = "400000000000000"
	MANAJEMEN_PENGGUNA_AKTIVASI_BUTTON_UPDATE                                                     Access = "800000000000000"
)
