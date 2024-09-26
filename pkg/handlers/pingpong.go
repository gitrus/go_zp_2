package handlers

import (
	"fmt"
	"net/http"
)

func PingPongHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Fprintln(w, "pong")
		return
	}

	http.Error(w, "Invalid request", http.StatusBadRequest)
}
