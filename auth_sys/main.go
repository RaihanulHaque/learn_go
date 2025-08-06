// package authsys
package main

import (
	"auth_sys/db"
	"auth_sys/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	db.InitDB()

    r := mux.NewRouter()

    r.HandleFunc("/signup", handlers.SignUp).Methods("POST")
    r.HandleFunc("/login", handlers.Login).Methods("POST")
    r.HandleFunc("/profile", handlers.Profile).Methods("GET")

    log.Println("Server started at :8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}
