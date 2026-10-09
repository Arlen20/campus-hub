package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
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

	for i:=0;i<len(words);i++ {
		w := words[i]
		if w == "(hex)" || w == "(bin)" {
			if len(result) > 0 {
				base := 16
				if w == "(bin)" {
					base = 2
				}
				last := len(result) - 1
				n, err := strconv.ParseInt(result[last], base, 64)
				if err == nil {                                     
					result[last] = strconv.FormatInt(n, 10)         
				}  
			}
			continue
		}
		if w == "(up)" || w == "(low)" || w == "(cap)" {
			if len(result) > 0 {
				mode := strings.Trim(w, "()")   // "(up)" → "up"
				last := len(result) - 1
				result[last] = transform(result[last], mode)
			}
			continue
		}
		if (w == "(up," || w == "(low," || w == "(cap,") && i+1 < len(words) {
			mode := strings.Trim(w, "(,")
			n,err := strconv.Atoi(strings.TrimSuffix(words[i+1], ")"))
			if err ==nil{
				for j:= len(result)-n;j<len(result);j++{
					result[j] = transform(result[j], mode)
				}
				i++
				continue
			}
		}
		result = append(result, w)
	}
	result = fixArticles(result)
	result = fixPunctuation(result)
	return strings.Join(result, " ")
}

func fixPunctuation(words []string) []string {
	out := []string{}

	for _, w := range words {
		// считаем, сколько знаков в начале слова
		k := 0
		for k < len(w) && strings.ContainsRune(".,!?:;", rune(w[k])) {
			k++
		}
		punct := w[:k] // знаки из начала, например ","
		rest := w[k:]  // остаток, например "and"

		if punct != "" && len(out) > 0 {
			out[len(out)-1] += punct
		} else {
			rest = w // приклеивать некуда или нечего — оставляем слово целиком
		}

		// TODO 2: если rest не пустой, добавь его в out
		if rest != "" {
			out = append(out, rest)
		}
	}

	return out
}

func fixArticles(words []string) []string {
	for i := 0; i < len(words); i++ {
		if (words[i] == "a" || words[i] == "A") && i+1 < len(words) {
			next := words[i+1]
			if strings.ContainsRune("aeiouhAEIOUH", rune(next[0])) {
				// TODO: если words[i] == "a", замени на "an"
				if words[i] == "a"{
					words[i] = "an"
				}else {
					words[i] = "An"
				}
				//       если words[i] == "A", замени на "An"
			}
		}
	}
	return words
}