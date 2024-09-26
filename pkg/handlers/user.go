package handlers

import (
	"encoding/json"
	"net/http"
)

type UserContactRequest struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Age   int    `json:"age"`
}

type UserContactResponse struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func User(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request", http.StatusMethodNotAllowed)
		return
	}

	var ucr UserContactRequest
	err := json.NewDecoder(r.Body).Decode(&ucr)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	var rucr UserContactResponse = UserContactResponse{
		Name: ucr.Name,
		Age:  ucr.Age,
	}
	err = json.NewEncoder(w).Encode(rucr)
	if err != nil {

	}
}
