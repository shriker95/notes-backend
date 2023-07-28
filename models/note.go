package models

type Note struct {
	ID   uint   `json:"id" gorm:"primary_key"`
	TEXT string `json:"text"`
}
