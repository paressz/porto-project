package model

type Auth struct {
	Token string `gorm:"type:text;not null;" json:"token"`
}