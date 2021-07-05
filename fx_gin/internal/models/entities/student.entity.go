package entities

type StudentEntity struct {
	ID int `json:"id" gorm:"primary_key"`
	Name string `json:"name" gorm:"not null"`
}
