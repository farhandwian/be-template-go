package model

type Status string

const (
	MenungguVerifikasi Status = "Menunggu Verifikasi"
	Ditolak            Status = "Ditolak"
	Disetujui          Status = "Disetujui"
)
