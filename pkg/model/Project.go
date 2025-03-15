package model

type Project struct {
	Id				string 	`gorm:"primaryKey" json:"id"`
	Name 			string 	`gorm:"size:50; not null" json:"name"`
	Description 	string 	`gorm:"text" json:"description"`
	ImageUrl 		string 	`gorm:"type:text;not null;" json:"imageUrl"`
	IntId			int		`gorm:"autoIncrement" json:"intId"`
}