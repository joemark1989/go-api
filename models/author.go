package models

type Author struct {
	Id        int    `json:"id" gorm:"primaryKey"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}
