package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct{
	Name  string `json:"name"`
	Email string `json:"email"`
	ID   string    `json:"id"`
}

var users = make(map[string]User) // It creates a map to store users. Make is used to initialize the map.


func getAllUser( w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	var userList []User
	for _, user := range users {
		userList = append(userList, user)
	}
	json.NewEncoder(w).Encode(userList)
}

func getUser(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	user, exists := users[params["id"]]
	fmt.Println(user)
	fmt.Println(exists)
	json.NewEncoder(w).Encode(user)
}

func createUser(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	var user User
	if err:= json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	users[user.ID] = user
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func main() {
	// Inintialize the router
	router := mux.NewRouter()

	// Define the endpoints
	router.HandleFunc("/users",getAllUser).Methods("GET")
	router.HandleFunc("/users/{id}", getUser).Methods("GET")
	router.HandleFunc("/users", createUser).Methods("POST")

	// Start server
	fmt.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
	
}