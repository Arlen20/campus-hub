package main

import (
	"fmt"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprint(w, "Hello from Campus Hub!")
}

func aboutpage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Это учебный проект Арлена")
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/about", aboutpage)

	fmt.Println("Сервер запущен: http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("ошибка:", err)
	}
}
