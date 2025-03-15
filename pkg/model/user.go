package model

type User struct {
	Name string `gorm:"not null" json:"name"`
	Username string `gorm:"type:varchar(12);not null; uniqueIndex" json:"username"`
	Email string `gorm:"primaryKey;uniqueIndex;not null" json:"email"`
	Password string `gorm:"not null" json:"password"`
}