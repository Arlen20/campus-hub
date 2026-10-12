package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"

	"github.com/Arlen20/campus-hub/go/internal/banner"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println(`usage: banner "text"`)
		return
	}
	text := os.Args[1]
	
	data, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Println("ошибка:", err)
		return
	}

	content := strings.ReplaceAll(string(data), "\r\n", "\n")
	lines := strings.Split(content, "\n")

	for _, ch := range text {
		if ch < 32 || ch > 126 {
			fmt.Printf("ошибка: символ %q не поддерживается\n", ch)
			return
		}	
	}

	fmt.Print(banner.Render(text, lines))

	analyze(text)
}

func analyze(text string) {
	clean := strings.ReplaceAll(text, "\\n", "")
	upper, lower, digits, special := 0, 0, 0, 0

	for _, ch := range clean {
		if unicode.IsUpper(ch) {
			upper++
		} else if unicode.IsLower(ch) {
			lower++
		} else if unicode.IsDigit(ch) {
			digits++
		} else{
			special++
		}
	}

	mixed := "no"

	if upper > 0 && lower > 0 {
		mixed = "yes"
	}

	complexity := lower*1 + upper*2 + digits*2 + special*3

	fmt.Println()
	fmt.Println("Uppercase:", upper)
	fmt.Println("Lowercase:", lower)
	fmt.Println("Digits:   ", digits)
	fmt.Println("Special:  ", special)
	fmt.Println("Mixed case:", mixed)
	fmt.Println("Complexity:", complexity)
}