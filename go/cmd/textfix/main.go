package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("usage: textfix <input> <output>")
		return
	}
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("ошибка:", err)
		return
	}
	text := fix(string(data))
	err = os.WriteFile(os.Args[2], []byte(text), 0644)
	if err != nil {
		fmt.Println("ошибка:", err)
		return
	}
}

func transform(word string, mode string) string {
	switch mode {
	case "up":
		// TODO: верни word большими буквами
		return strings.ToUpper(word)
	case "low":
		// TODO: верни word маленькими буквами
		return strings.ToLower(word)
	case "cap":
		// TODO: верни word с большой первой буквой
		return strings.ToUpper(word[:1]) + strings.ToLower(word[1:])
	}
	return word
}

func fix(text string) string {
	words := strings.Fields(text)
	result := []string{}

	for _, w := range words {
		if w == "(up)" || w == "(low)" || w == "(cap)" {
			if len(result) > 0 {
				mode := strings.Trim(w, "()")   // "(up)" → "up"
				last := len(result) - 1
				result[last] = transform(result[last], mode)
			}
			continue
		}
		result = append(result, w)
	}

	return strings.Join(result, " ")
}