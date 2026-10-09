package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
	"sort"
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
	fmt.Println("Language:", detectLanguage(text))
	fmt.Println("Keywords:", strings.Join(topKeywords(text, 5), ", "))
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
					if j >= 0 {
						result[j] = transform(result[j], mode)
					}
				}
				i++
				continue
			}
		}
		result = append(result, w)
	}
	result = fixArticles(result)
	result = fixPunctuation(result)
	result = fixQuotes(result)
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

func fixQuotes(words []string) []string {
	out := []string{}
	open := false

	for i := 0; i < len(words); i++ {
		w := words[i]

		if w == "'" {
			if !open && i+1 < len(words) {
				// открывающая: приклеиваем к следующему слову
				words[i+1] = "'" + words[i+1]
				open = true
				continue
			}
			if open && len(out) > 0 {
				// TODO 1: закрывающая — приклей "'" к последнему слову в out
				out[len(out)-1] += "'"
				open = false
				// TODO 2: поставь флажок обратно в false
				continue
			}
		}

		out = append(out, w)
	}

	return out
}

var englishWords = map[string]bool{
	"the": true, "and": true, "is": true, "of": true, "to": true,
	"in": true, "it": true, "you": true, "that": true, "with": true,
}

var frenchWords = map[string]bool{
	"le": true, "la": true, "les": true, "et": true, "est": true,
	"un": true, "une": true, "de": true, "des": true, "je": true,
}

func detectLanguage(text string) string {
	en := 0
	fr := 0

	// 1. считаем частые слова
	for _, w := range strings.Fields(strings.ToLower(text)) {
		w = strings.Trim(w, ".,!?:;'")
		if englishWords[w] == true {
			en++
		}
		if frenchWords[w] == true {
			fr++
		}
	}

	// 2. считаем буквы с акцентами
	for _, r := range text {
		if strings.ContainsRune("éèêàçùôîâëï", r) {
			fr++
		}
	}

	// TODO 2:
	// если en == 0 и fr == 0 → верни "Unknown"
	if en == 0 && fr == 0 {
		return "Unknown"
	}
	if fr > en {
		return "French"
	}
	return "English"
	
	// если fr > en           → верни "French"
	// иначе                  → верни "English"
	return ""
}

var stopWords = map[string]bool{
	"the": true, "a": true, "an": true, "is": true, "are": true, "was": true,
	"and": true, "or": true, "of": true, "to": true, "in": true, "on": true,
	"it": true, "this": true, "that": true, "i": true, "you": true, "he": true,
	"she": true, "we": true, "they": true, "with": true, "for": true,
	"le": true, "la": true, "les": true, "et": true, "est": true, "un": true,
	"une": true, "de": true, "des": true, "je": true, "il": true, "elle": true,
	"sur": true, "très": true,
}

func topKeywords(text string, n int) []string {
	counts := map[string]int{}

	for _, w := range strings.Fields(strings.ToLower(text)) {
		w = strings.Trim(w, ".,!?:;'")
		if w == "" || stopWords[w] { // пустое или мусорное — пропускаем
			continue
		}
		counts[w]++ // считаем слово
	}

	keys := []string{}
	for w := range counts {
		keys = append(keys, w)
	}

	sort.Slice(keys, func(i, j int) bool {
		if counts[keys[i]] != counts[keys[j]] {
			return counts[keys[i]] > counts[keys[j]]
		}
		return keys[i] < keys[j]
	})

	if len(keys) > n { // слов больше, чем нужно — обрезаем
		keys = keys[:n]
	}
	return keys
}