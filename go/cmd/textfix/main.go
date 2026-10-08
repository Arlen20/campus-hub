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

func fix(text string) string {
	words := strings.Fields(text)
	result := []string{}

	for _, w := range words {
		if w == "(up)" {
			// TODO 1: если len(result) > 0,
			if len(result) > 0 {
				last := len(result) - 1
				result[last] = strings.ToUpper(result[last])
			}
			continue
		}
		result = append(result, w)
	}
	return strings.Join(result, " ")
}
