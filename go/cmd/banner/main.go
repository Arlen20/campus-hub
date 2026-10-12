package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
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

	parts := strings.Split(text, "\\n")

	for _, ch := range text {
		if ch < 32 || ch > 126 {
			fmt.Printf("ошибка: символ %q не поддерживается\n", ch)
			return
		}	
	}

	for _, part := range parts {
		if part == "" {
			fmt.Println()
		}else{
			printBanner(part,lines)
		}
	}
	analyze(text)
}


func printBanner(word string, lines []string){
	for row := 0; row < 8; row++ {          // для каждой из 8 строк
		line := ""
		for _, ch := range word {           // пройти по всем буквам
			start := (int(ch)-32)*9 + 1
			line += lines[start+row]        // приклеить кусок этой буквы
		}
		fmt.Println(line)                   // напечатать собранную строку
	}
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