package student

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type StudentController struct {
	STUDENTS []Student
}

func NewStudentController(router *mux.Router) *StudentController {
	students := []Student{
		{ID: 1, NAME: "name1", EMAIL: "name1@domain"},
		{ID: 2, NAME: "name2", EMAIL: "name2@domain"},
		{ID: 3, NAME: "name3", EMAIL: "name3@domain"},
	}

	sc := &StudentController{students}

	usersR := router.PathPrefix("/students").Subrouter()
	usersR.Path("").Methods(http.MethodGet).HandlerFunc(sc.getAllStudents)
	usersR.Path("").Methods(http.MethodPost).HandlerFunc(sc.saveStudent)
	usersR.Path("/{id}").Methods(http.MethodGet).HandlerFunc(sc.getStudent)
	usersR.Path("/{id}").Methods(http.MethodPut).HandlerFunc(sc.updateStudent)
	usersR.Path("/{id}").Methods(http.MethodDelete).HandlerFunc(sc.deleteStudent)
	return sc
}

func (StudentController *StudentController) getAllStudents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(StudentController.STUDENTS)
}

func (sc *StudentController) getStudent(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for _, student := range sc.STUDENTS {
		if student.ID == id {
			json.NewEncoder(w).Encode(student)
			return
		}
	}
	http.Error(w, "User not found", http.StatusBadRequest)
}

func (sc *StudentController) saveStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var student Student
	json.NewDecoder(r.Body).Decode(&student)
	sc.STUDENTS = append(sc.STUDENTS, student)
	json.NewEncoder(w).Encode(MESSAGE{"User saved successfully"})
}

func (sc *StudentController) updateStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var updatedStudent Student
	json.NewDecoder(r.Body).Decode(&updatedStudent)
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	for index, student := range sc.STUDENTS {
		if student.ID == id {
			sc.STUDENTS = append(sc.STUDENTS[:index], sc.STUDENTS[index+1:]...)
			sc.STUDENTS = append(sc.STUDENTS, updatedStudent)
			json.NewEncoder(w).Encode(MESSAGE{"User updated successfully"})
			return
		}
	}
	http.Error(w, "User not found", http.StatusBadRequest)
}

func (sc *StudentController) deleteStudent(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	id, _ := strconv.Atoi(params["id"])

	for index, student := range sc.STUDENTS {
		if student.ID == id {
			sc.STUDENTS = append(sc.STUDENTS[:index], sc.STUDENTS[index+1:]...)
			json.NewEncoder(w).Encode(MESSAGE{"User deleted successfully"})
			return
		}
	}
	http.Error(w, "User not found", http.StatusBadRequest)
}
