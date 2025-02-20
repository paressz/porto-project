package models


type Project struct {
	Id				string 	`gorm:"primaryKey"`
	Name 			string 	`gorm:"size:50"`
	Description 	string 	`gorm:"text"`
	ImageUrl 		string 	`gorm:"text"`
	IntId			int		`gorm:"autoIncrement"`
}