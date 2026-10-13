package llm

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

const model = "gemini-flash-lite-latest"

type part struct {
	Text string `json:"text"`
}
type content struct {
	Parts []part `json:"parts"`
}
type request struct {
	Contents []content `json:"contents"`
}
type response struct {
	Candidates []struct {
		Content content `json:"content"`
	} `json:"candidates"`
}

func Ask(prompt string) (string, error) {
	key := os.Getenv("GEMINI_API_KEY")
	if key == "" {
		return "", errors.New("GEMINI_API_KEY не задан")
	}

	// коробка → текст JSON (как json.Marshal(p) в опыте 1)
	body, err := json.Marshal(request{
		Contents: []content{{Parts: []part{{Text: prompt}}}},
	})
	if err != nil {
		return "", err
	}

	// отправляем текст в Google
	url := "https://generativelanguage.googleapis.com/v1beta/models/" + model + ":generateContent"
	req, err := http.NewRequest("POST", url, bytes.NewReader(body))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-goog-api-key", key)

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("gemini: %s: %s", resp.Status, data)
	}

	// текст JSON → коробка (как json.Unmarshal в опыте 3)
	var r response
	if err := json.Unmarshal(data, &r); err != nil {
		return "", err
	}
	if len(r.Candidates) == 0 || len(r.Candidates[0].Content.Parts) == 0 {
		return "", errors.New("gemini: пустой ответ")
	}
	// достаём текст из матрёшки (как p2.Name, только глубже)
	return r.Candidates[0].Content.Parts[0].Text, nil
}
