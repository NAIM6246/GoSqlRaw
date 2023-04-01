package models

type DepartmentCreateDto struct {
	DepartmentName string `json:"departmentName"`
	DepartmentCode string `json:"departmentCode"`
}

type CostOfDepartment struct {
	ID             int64  `json:"id"`
	DepartmentName string `json:"departmentName"`
	TotalCost      int64  `json:"totalCost"`
}

type StudentsOfDepartment struct {
	ID             int64  `json:"id"`
	DepartmentName string `json:"departmentName"`
	TotalStudent   int64  `json:"totalStudent"`
}
