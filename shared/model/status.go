package model

type Status string

const (
	StatusMenungguVerifikasi Status = "Menunggu Verifikasi"
	StatusDitolak            Status = "Ditolak"
	StatusDisetujui          Status = "Disetujui"
)
