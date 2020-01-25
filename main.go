package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/users", getStudents).Methods("GET")
	router.HandleFunc("/users/add", addStudent).Methods("POST")
	router.HandleFunc("/users/update", updateUser).Methods("POST")
	router.HandleFunc("/users/delete", deleteUser).Methods("POST")

	http.Handle("/", router)
	fmt.Println("Connected to port 1234")
	log.Fatal(http.ListenAndServe(":1234", router))

}
