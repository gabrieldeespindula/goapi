package models

type Record struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}
