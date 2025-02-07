package model

type FCMToken struct {
	ID    int `gorm:"primaryKey"`
	Token string
}
