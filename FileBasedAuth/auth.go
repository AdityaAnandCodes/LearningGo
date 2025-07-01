package main

import (
	"net/http"
	"fmt"
	"os"
	"encoding/json"
	"io/ioutil"
)

type LoginData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserData struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Message string `json:"message"`
	Result  int    `json:"result"`
}

var fileName string = "data.json"

func handleUserData(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost{
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}
	var userData UserData
	err := json.NewDecoder(r.Body).Decode(&userData)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0644)
    if err != nil {
        panic(err)
    }
    defer file.Close()

	var users[] UserData

	data , _ := ioutil.ReadAll(file)
	if len(data) > 0 {
		err = json.Unmarshal(data, &users)
		if err != nil {
			http.Error(w, "Error reading data", http.StatusInternalServerError)
			return
		}
	}

	users = append(users, userData)

	file.Truncate(0)
	file.Seek(0,0)
	updatedUsers, err := json.Marshal(users)
	if err != nil {
		http.Error(w, "Error encoding data", http.StatusInternalServerError)
		return
	}
	_, err = file.Write(updatedUsers)
	if err != nil {
		http.Error(w, "Error writing data", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := LoginResponse{	
		Message: "User data saved successfully",
		Result:  1,
	}
	json.NewEncoder(w).Encode(response)

}


func handleLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method !=  http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}
	var loginData LoginData
	err:= json.NewDecoder(r.Body).Decode(&loginData)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	file, err := os.Open(fileName)
	if err != nil {
		http.Error(w, "Error opening file", http.StatusInternalServerError)
		return
	}
	defer file.Close()
	var users [] UserData
	data, err := ioutil.ReadAll(file)
	if err != nil {
		http.Error(w, "Error reading file", http.StatusInternalServerError)
		return
	}
	if len(data) > 0 {
		err = json.Unmarshal(data, &users)
		if err != nil {
			http.Error(w, "Error parsing data", http.StatusInternalServerError)
			return
		}
	}

	for _, user := range users {
		if user.Username == loginData.Username && user.Password == loginData.Password {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			response := LoginResponse{
				Message: "Login successful",
				Result:  2,
			}
			json.NewEncoder(w).Encode(response)
			return
		}
	} 
}


func main() {
	http.HandleFunc("/create-user", handleUserData)
	http.HandleFunc("/login", handleLogin)
	fmt.Println("Server is running on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic("Error starting server: " + err.Error())
	}

}