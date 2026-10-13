package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type PageData struct {
	Title string
}

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}
	data := PageData{Title: "Symbol Art Web"}
	tmpl.Execute(w, data)
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
