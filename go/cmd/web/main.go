package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"

	"github.com/Arlen20/campus-hub/go/internal/banner"
)

type PageData struct {
	Title  string
	Banner string
}

func asciiArt(w http.ResponseWriter, r *http.Request) {
	// 1. только POST
	if r.Method != http.MethodPost {
		http.Error(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// 2. текст из формы
	text := r.FormValue("text")
	text = strings.ReplaceAll(text, "\r\n", "\\n")

	if text == "" {
		http.Error(w, "400 Bad Request: empty text", http.StatusBadRequest)
		return
	}

	// 3. проверка символов
	for _, ch := range text {
		if ch < 32 || ch > 126 {
			http.Error(w, "400 Bad Request", http.StatusBadRequest)
			return
		}
	}

	// 4. шаблон баннеров
	data, err := os.ReadFile("standard.txt")
	if err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}
	content := strings.ReplaceAll(string(data), "\r\n", "\n")
	lines := strings.Split(content, "\n")

	// 5. рисуем
	result := banner.Render(text, lines)

	// 6. показываем страницу с баннером
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, PageData{Title: "Symbol Art Web", Banner: result})
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
	http.HandleFunc("/ascii-art", asciiArt)

	fmt.Println("Сервер запущен: http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("ошибка:", err)
	}
}
