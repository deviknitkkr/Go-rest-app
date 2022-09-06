package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type Student struct {
	ID    int    `json:"id"`
	NAME  string `json:"name"`
	EMAIL string `json:"email"`
}

type MESSAGE struct {
	MESSAGE string `json:"message"`
}

var students []Student

func main() {

	students = append(students, Student{ID: 1, NAME: "name1", EMAIL: "name1@mail.com"})
	students = append(students, Student{ID: 2, NAME: "name2", EMAIL: "name2@mail.com"})
	students = append(students, Student{ID: 3, NAME: "name3", EMAIL: "name3@mail.com"})

	router := mux.NewRouter()

	router.HandleFunc("/students", getAllStudents).Methods("GET")
	router.HandleFunc("/students/{id}", getStudent).Methods("GET")
	router.HandleFunc("/students", saveStudent).Methods("POST")
	router.HandleFunc("/students/{id}", updateStudent).Methods("PUT")
	router.HandleFunc("/students/{id}", deleteStudent).Methods("DELETE")

	fmt.Println("Starting server...")
	http.ListenAndServe(":8080", router)
}

func getAllStudents(w http.ResponseWriter, r *http.Request) {
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
