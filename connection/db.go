package connection

import (
	"fmt"
	"goSqlRaw/models"
)

type QueryInterface interface {
	CalculateCostOfDepartments() ([]*models.CostOfDepartment, error)
	HighestSalaryGettingTwoTeacher() ([]*models.TeachersDetails, error)
	TotalStudentOfEacheDepartment() ([]*models.StudentsOfDepartment, error)
}

func GetQuerInterface() QueryInterface {
	return dbInstance
}

func (db *DB) CalculateCostOfDepartments() ([]*models.CostOfDepartment, error) {
	var cost []*models.CostOfDepartment

	err := db.Raw(`SELECT D.ID AS ID, D.DEPARTMENT_NAME AS DEPARTMENT_NAME, SUM(T.SALARY) as TOTAL_COST
					FROM DEPARTMENTS AS D INNER JOIN TEACHERS AS T
					ON D.ID=T.DEPARTMENT_ID
					GROUP BY D.ID, D.DEPARTMENT_NAME
					ORDER BY TOTAL_COST DESC`).Scan(&cost)
	if err.Error != nil {
		return nil, err.Error
	}
	fmt.Println(cost)
	return cost, nil
}

func (db *DB) HighestSalaryGettingTwoTeacher() ([]*models.TeachersDetails, error) {
	var teachersDetails []*models.TeachersDetails

	err := db.Raw(`SELECT T.ID AS ID, D.DEPARTMENT_NAME AS DEPARTMENT_NAME, T.TEACHER_NAME AS TEACHER_NAME, T.SALARY AS SALARY
					FROM TEACHERS AS T
					INNER JOIN DEPARTMENTS AS D
					ON T.DEPARTMENT_ID = D.ID
					ORDER BY SALARY DESC
					LIMIT 2`).Scan(&teachersDetails)
	if err.Error != nil {
		return nil, err.Error
	}
	fmt.Println(teachersDetails)
	return teachersDetails, nil
}

func (db *DB) TotalStudentOfEacheDepartment() ([]*models.StudentsOfDepartment, error) {
	var numberOfStudentPerDepartent []*models.StudentsOfDepartment

	err := db.Raw(`SELECT D.ID AS ID, D.DEPARTMENT_NAME AS DEPARTMENT_NAME, COUNT(S.ID) AS TOTAL_STUDENT
					FROM DEPARTMENTS AS D 
					INNER JOIN TEACHERS AS T ON D.ID = T.DEPARTMENT_ID
					INNER JOIN STUDENTS AS S ON T.ID = S.TEACHER_ID
					GROUP BY D.ID, D.DEPARTMENT_NAME
					ORDER BY TOTAL_STUDENT DESC`).Scan(&numberOfStudentPerDepartent)

	if err.Error != nil {
		return nil, err.Error
	}
	fmt.Println(numberOfStudentPerDepartent)
	return numberOfStudentPerDepartent, nil
}

func (db *DB) AddStudent(data *models.StudentCreateDto) error {
	return db.Exec(`INSERT INTO STUDENTS (STUDENT_NAME, TEACHER_ID) VALUES (?, ?)`, data.StudentName, data.TeacherID).Error
}

func (db *DB) AddTeacher(data *models.TeacherCreateDto) error {
	return db.Exec(`INSERT INTO TEACHERS (TEACHER_NAME, SALARY, DEPARTMENT_ID) VALUES (?, ?, ?)`, data.TeacherName, data.Salary, data.DepartmentID).Error
}

func (db *DB) AddDepartment(data *models.DepartmentCreateDto) error {
	return db.Exec(`INSERT INTO DEPARTMENTS (DEPARTMENT_NAME, DEPARTMENT_CODE) VALUES  (?, ?)`, data.DepartmentName, data.DepartmentCode).Error
}
