package main

import (
	"fmt"
	"net/http"
	"os"
)

func artists(w http.ResponseWriter, r *http.Request) {
	// 1. читаем готовый JSON из файла
	data, err := os.ReadFile("data/artists.json")
	if err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}

	// 2. говорим клиенту: «внутри JSON»
	w.Header().Set("Content-Type", "application/json")

	// 3. отдаём как есть
	w.Write(data)
}

func main() {
	http.HandleFunc("/api/artists", artists)

	fmt.Println("Fake API: http://localhost:8081/api/artists")
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		fmt.Println("ошибка:", err)
	}
}
