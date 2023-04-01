package handlers

import (
	"encoding/json"
	"goSqlRaw/connection"
	"goSqlRaw/models"
	"goSqlRaw/utils"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	db     *connection.DB
	writer *utils.CsvWriter
}

type IHandler interface {
	Handle(router chi.Router)
}

func NewHandler() IHandler {
	return &Handler{
		db:     connection.GetDBInstance(),
		writer: utils.GetCSVFileWriter(),
	}
}

func (h *Handler) Handle(router chi.Router) {
	router.Get("/calculate-department-costs", h.calculateDepartmentCost)
	router.Get("/highest-salary-teachers", h.highestSalaryGettingTeachers)
	router.Get("/departments-total-student", h.numOfStudentPerDepartment)
	router.Post("/add-student", h.addStudents)
	router.Post("/add-teacher", h.addTeacher)
	router.Post("/add-department", h.addDepartment)
}

func (h *Handler) calculateDepartmentCost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	costs, err := h.db.CalculateCostOfDepartments()
	if err != nil {
		h.writer.Write(err, "calculateDepartmentCost", "error while getting data from db")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message" : "internal server error"}`))
		return
	}

	h.writer.Write(costs, "calculateDepartmentCost", "success")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(costs)
}

func (h *Handler) highestSalaryGettingTeachers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	teachers, err := h.db.HighestSalaryGettingTwoTeacher()
	if err != nil {
		h.writer.Write(err, "highestSalaryGettingTeachers", "error while getting data from db")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message" : "internal server error"}`))
		return
	}

	h.writer.Write(teachers, "highestSalaryGettingTeachers", "success")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(teachers)
}

func (h *Handler) numOfStudentPerDepartment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	departmentsStudent, err := h.db.TotalStudentOfEacheDepartment()
	if err != nil {
		h.writer.Write(err, "numOfStudentPerDepartment", "error while getting data from db")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message" : "internal server error"}`))
		return
	}

	h.writer.Write(departmentsStudent, "numOfStudentPerDepartment", "success")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(departmentsStudent)
}

func (h *Handler) addStudents(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	var payload *models.StudentCreateDto
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		h.writer.Write(err, "addStudents", "error while parsing data")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message" : "failed to parse request body"}`))
		return
	}

	err := h.db.AddStudent(payload)
	if err != nil {
		h.writer.Write(err, "addStudents", "error while inserting data")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message" : "failed to add student}`))
		return
	}

	h.writer.Write(payload, "addStudents", "success")
	w.WriteHeader(http.StatusCreated)
	log.Println("student created", "data", payload)
}

func (h *Handler) addTeacher(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	var payload *models.TeacherCreateDto
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		h.writer.Write(err, "addTeacher", "error while parsing data")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message" : "failed to parse request body}`))
		return
	}

	err := h.db.AddTeacher(payload)
	if err != nil {
		h.writer.Write(err, "addTeacher", "error while inserting data")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message" : "failed to add teacher}`))
		return
	}

	h.writer.Write(payload, "addTeacher", "success")
	w.WriteHeader(http.StatusCreated)
	log.Println("teacher created", "data", payload)
}

func (h *Handler) addDepartment(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	var payload *models.DepartmentCreateDto
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		h.writer.Write(err, "addDepartment", "error while parsing data")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message" : "failed to parse request body}`))
		return
	}

	err := h.db.AddDepartment(payload)
	if err != nil {
		h.writer.Write(err, "addDepartment", "error while inserting data")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message" : "failed to add department}`))
		return
	}

	h.writer.Write(payload, "addDepartment", "success")
	w.WriteHeader(http.StatusCreated)
	log.Println("department created", "data", payload)
}
