package request

type Students struct {
	StudentID int64  `form:"student_id"`
	FullName  string `form:"full_name"`
	ExtraID   int64  `form:"extra_id"`
}