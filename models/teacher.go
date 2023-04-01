package models

type TeacherCreateDto struct {
	TeacherName  string `json:"teacherName"`
	Salary       int    `json:"salary"`
	DepartmentID int    `json:"departmentID"`
}


type TeachersDetails struct {
	ID           int    `json:"id"`
	TeacherName  string `json:"teacherName"`
	Salary       int    `json:"salary"`
	DepartmentName string    `json:"departmentName"`
}