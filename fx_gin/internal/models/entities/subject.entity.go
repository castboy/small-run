package entities

type SubjectEntity struct {
	ID   int    `json:"id" gorm:"primary_key"`
	Name string `json:"name" gorm:"not null"`
}
