package models

type StudentCreateDto struct {
	StudentName string `json:"studentName"`
	TeacherID   int    `json:"teacherID"`
}
