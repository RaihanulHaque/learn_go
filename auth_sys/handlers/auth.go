package handlers

import (
	"auth_sys/db"
	"auth_sys/utils"
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type Credentials struct{
	Username string "json:username"
	Password string "json:password"
}

func SignUp(w http.ResponseWriter, r *http.Request){
	var cred Credentials
	err := json.NewDecoder(r.Body).Decode(&cred)
	if err != nil || cred.Username =="" || cred.Password == ""{
		http.Error(w, "Invalid Input", http.StatusBadRequest)
		return
	}

	hashed_pass, err := bcrypt.GenerateFromPassword([]byte(cred.Password), bcrypt.DefaultCost)
	fmt.Println("Hashed Password", hashed_pass)
	if err != nil {
        http.Error(w, "Error hashing password", http.StatusInternalServerError)
        return
    }

	_, err = db.DB.Exec("INSERT INTO users(username, password) VALUES (?, ?)", cred.Username, string(hashed_pass))
    if err != nil {
        http.Error(w, "Username already exists", http.StatusConflict)
        return
    }

    w.WriteHeader(http.StatusCreated)
    w.Write([]byte("User created successfully"))
}

func Login(w http.ResponseWriter, r *http.Request) {
    var creds Credentials
    err := json.NewDecoder(r.Body).Decode(&creds)
    if err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    var storedHashedPassword string
    row := db.DB.QueryRow("SELECT password FROM users WHERE username = ?", creds.Username)
    err = row.Scan(&storedHashedPassword)
    if err != nil {
        http.Error(w, "User not found", http.StatusUnauthorized)
        return
    }

    err = bcrypt.CompareHashAndPassword([]byte(storedHashedPassword), []byte(creds.Password))
    if err != nil {
        http.Error(w, "Wrong password", http.StatusUnauthorized)
        return
    }

    token, err := utils.GenerateJWT(creds.Username)
    if err != nil {
        http.Error(w, "Error generating token", http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(map[string]string{"token": token})
}


func Profile(w http.ResponseWriter, r *http.Request) {
    authHeader := r.Header.Get("Authorization")
    if authHeader == "" {
        http.Error(w, "Missing token", http.StatusUnauthorized)
        return
    }

    tokenStr := authHeader[len("Bearer "):] // Strip "Bearer "

    username, err := utils.ValidateJWT(tokenStr)
    if err != nil {
        http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
        return
    }

    json.NewEncoder(w).Encode(map[string]string{
        "message": "Welcome, " + username,
    })
}


