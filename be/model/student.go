package model

type Students struct {
	StudentID int64  `json:"student_id"`
	FullName  string `json:"full_name"`
	ExtraName string `json:"extra_name" gorm:"column:extra_name"`
}