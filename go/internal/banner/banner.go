package banner

import "strings"

// Render возвращает готовый баннер одной строкой.
func Render(text string, lines []string) string {
	result := ""
	parts := strings.Split(text, "\\n")

	for _, part := range parts {
		if part == "" {
			result += "\n"
			continue
		}
		for row := 0; row < 8; row++ {
			line := ""
			for _, ch := range part {
				start := (int(ch)-32)*9 + 1
				line += lines[start+row]
			}
			result += line + "\n"
		}
	}
	return result
}