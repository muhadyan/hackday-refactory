package model

// import (
// 	"github.com/jinzhu/gorm"
// )

type Students struct {
	StudentID int64  `json:"student_id" gorm:"not null;primaryKey;autoIncrement"`
	FullName  string `json:"full_name"`
	ExtraName string `json:"extra_name" gorm:"column:extra_name"`
}

// for post
type students struct {
	StudentID int64  `form:"student_id"`
	FullName  string `form:"full_name"`
	ExtraID   int64  `form:"extra_id"`
}