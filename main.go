package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"time"
	"io"
	"fmt"
	"strconv"
)

func handler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello, world!\n")
}

// Route declaration
func router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", handler)
	return r
}

func startRouter(port int){
	router := router()
        srv := &http.Server{
                Handler: router,
                Addr:    "127.0.0.1:"+strconv.Itoa(port),
                WriteTimeout: 15 * time.Second,
                ReadTimeout:  15 * time.Second,
        }
	fmt.Println("Starting Server at:",port)
        srv.ListenAndServe()
}

// Initiate web server
func main() {
	startRouter(8080)
}
