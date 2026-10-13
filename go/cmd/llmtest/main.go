package main

import (
	"fmt"

	"github.com/Arlen20/campus-hub/go/internal/llm"
)

func main() {
	answer, err := llm.Ask("Придумай 3 коротких слогана для IT-школы, каждый не длиннее 4 слов, на английском")
	if err != nil {
		fmt.Println("ошибка:", err)
		return
	}
	fmt.Println(answer)
}
