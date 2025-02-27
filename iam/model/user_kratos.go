package model

import (
	"time"
)

type UserKratosID string

type UserKratosCreate struct {
	Email         string `json:"email"`
	Password      string `json:"password"`
	Nama          string `json:"nama"`
	NoTelepon     string `json:"no_telepon"`
	Jabatan       string `json:"jabatan"`
	AksesPengguna string `json:"akses_pengguna"`
	JenisKelamin  string `json:"jenis_kelamin"`
}

type UserKratosGet struct {
	ID            UserKratosID `json:"id,omitempty"`
	Email         string       `json:"email"`
	Password      string       `json:"password"`
	Nama          string       `json:"nama"`
	NoTelepon     string       `json:"no_telepon"`
	Jabatan       string       `json:"jabatan"`
	AksesPengguna string       `json:"akses_pengguna"`
	JenisKelamin  string       `json:"jenis_kelamin"`
	CreatedAt     time.Time    `json:"created_at,omitempty"`
	UpdatedAt     time.Time    `json:"updated_at,omitempty"`
}

type UserKratosUpdate struct {
	ID            UserKratosID `json:"id,omitempty"`
	Email         string       `json:"email"`
	Password      string       `json:"password"`
	Nama          string       `json:"nama"`
	NoTelepon     string       `json:"no_telepon"`
	Jabatan       string       `json:"jabatan"`
	AksesPengguna string       `json:"akses_pengguna"`
	JenisKelamin  string       `json:"jenis_kelamin"`
}
