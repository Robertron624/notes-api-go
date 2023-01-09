package models

type Note struct {
	Id          int    `json:"id" gorm:"primaryKey"`
	Title       string `json:"title"`
	username    string `json:"username"`
	Description string `json:"description"`
}
