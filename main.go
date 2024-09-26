package main

import (
	"fmt"
	"go_zp_2/pkg/server_6_1"
	"net/http"
)

func main() {
	http.HandleFunc(`/`, server_6_1.MainPageHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("ошибка запуска сервера: $s\n", err.Error())
		return
	}
}
