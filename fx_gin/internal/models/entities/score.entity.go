package entities

type ScoreEntity struct {
	ID int `json:"id" gorm:"primary_key"`
	StudentID int `json:"student_id" gorm:"not null"`
	SubjectID int `json:"subject_id" gorm:"not null"`
	Score int `json:"score" gorm:"not null"`
}
