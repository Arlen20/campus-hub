package main

import (
	"fmt"
	"os"
	"strings"
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