package student

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

var students []Student

func NewStudentController(router *mux.Router) {
	students = []Student{
		{ID: 1, NAME: "name1", EMAIL: "name1@mail.com"},
		{ID: 2, NAME: "name2", EMAIL: "name2@mail.com"},
		{ID: 3, NAME: "name3", EMAIL: "name3@mail.com"},
	}

	router.HandleFunc("/students", getAllstudents).Methods("GET")
	router.HandleFunc("/students/{id}", getStudent).Methods("GET")
	router.HandleFunc("/students", saveStudent).Methods("POST")
	router.HandleFunc("/students/{id}", updateStudent).Methods("PUT")
	router.HandleFunc("/students/{id}", deleteStudent).Methods("DELETE")
}

func getAllstudents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(students)
	if err != nil {
		return
	}
}

func getStudent(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for _, student := range students {
		if student.ID == id {
			json.NewEncoder(w).Encode(student)
			return
		}
	}
	json.NewEncoder(w).Encode(MESSAGE{"User not found"})
}

func saveStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var student Student
	json.NewDecoder(r.Body).Decode(&student)
	students = append(students, student)
	json.NewEncoder(w).Encode(MESSAGE{"User saved successfully"})
}

func updateStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var updatedStudent Student
	json.NewDecoder(r.Body).Decode(&updatedStudent)
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	for index, student := range students {
		if student.ID == id {
			students = append(students[:index], students[index+1:]...)
			students = append(students, updatedStudent)
			json.NewEncoder(w).Encode(MESSAGE{"User updated successfully"})
			return
		}
	}
	json.NewEncoder(w).Encode(MESSAGE{"User not found"})
}

func deleteStudent(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	id, _ := strconv.Atoi(params["id"])

	for index, student := range students {
		if student.ID == id {
			students = append(students[:index], students[index+1:]...)
			json.NewEncoder(w).Encode(MESSAGE{"User deleted successfully"})
			return
		}
	}
	json.NewEncoder(w).Encode(MESSAGE{"User not found"})
}
