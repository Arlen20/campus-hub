package tour

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const baseURL = "http://localhost:8081/api"

type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}

var client = &http.Client{Timeout: 10 * time.Second}

// GetArtists забирает список всех групп.
func GetArtists() ([]Artist, error) {
	resp, err := client.Get(baseURL + "/artists")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("api: %s", resp.Status)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var artists []Artist
	err = json.Unmarshal(data, &artists)

	if err != nil {
		return nil, err
	}
	return artists, nil
}
