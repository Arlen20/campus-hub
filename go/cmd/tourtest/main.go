package main

import (
	"fmt"

	"github.com/Arlen20/campus-hub/go/internal/tour"
)

func main() {
	artists, err := tour.GetArtists()
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("all groups", len(artists))

	for _, a := range artists {
		fmt.Println(a.ID, a.Name, a.CreationDate)
	}
}
