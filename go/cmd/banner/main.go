package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Println("ошибка:", err)
		return
	}

	content := strings.ReplaceAll(string(data), "\r\n", "\n")
	lines := strings.Split(content, "\n")

	ch := '!'

	// TODO 1: посчитай start по формуле
	start := (int(ch) - 32) * 9 + 1
	for i := 0; i < 8; i++{
		fmt.Println(lines[start+i])
	}

	// TODO 2: напечатай 8 строк, начиная с lines[start]
	// подсказка: цикл for i := 0; i < 8; i++ и fmt.Println(lines[start+i])
}