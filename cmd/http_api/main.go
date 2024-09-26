package main

import (
	"fmt"
	"go_zp_2/pkg/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/ping", handlers.PingPongHandler)
	http.HandleFunc("/user", handlers.User)

	fmt.Println("Starting server on :1212...")
	if err := http.ListenAndServe(":1212", nil); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
