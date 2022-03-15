package request

type Students struct {
	StudentID int64  `json:"student_id"`
	FullName  string `json:"full_name"`
	ExtraID   int64  `json:"extra_id"`
}