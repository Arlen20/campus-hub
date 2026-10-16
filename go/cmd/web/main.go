package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"

	"github.com/Arlen20/campus-hub/go/internal/banner"
	"github.com/Arlen20/campus-hub/go/internal/llm"
)

type PageData struct {
	Title       string
	Banner      string
	Suggestions string
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

func suggest(w http.ResponseWriter, r *http.Request) {
	// 1. только POST
	if r.Method != http.MethodPost {
		http.Error(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// 2. тема из формы
	theme := strings.TrimSpace(r.FormValue("theme"))
	if theme == "" {
		http.Error(w, "400 Bad Request: empty theme", http.StatusBadRequest)
		return
	}

	// 3. строгий промпт: только то, что умеет рисовать баннер
	prompt := "Suggest 3 short texts for an ASCII art banner on the theme: " + theme + ". " +
		"Rules: English only, uppercase letters, digits and spaces only, " +
		"max 12 characters each, one per line, no numbering, no explanations, no markdown."

	// 4. спрашиваем ИИ
	answer, err := llm.Ask(prompt)
	if err != nil {
		answer = "ИИ сейчас недоступен, попробуй позже"
	}

	// 5. показываем страницу с идеями
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, PageData{Title: "Symbol Art Web", Suggestions: answer})
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
	http.HandleFunc("/suggest", suggest)

	fmt.Println("Сервер запущен: http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("ошибка:", err)
	}
}
