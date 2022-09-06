package main

import (
	"fmt"
	"github.com/deviknitkkr/Go-rest-app/student"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	student.NewStudentController(router)

	fmt.Println("Starting server...")
	http.ListenAndServe(":8080", router)
}
