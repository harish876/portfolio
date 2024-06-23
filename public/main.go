package main

import (
	"context"
	"log"
	"os"

	"github.com/harish876/portfolio/views/about"
)

func main() {
	f, err := os.Create("about.html")
	if err != nil {
		log.Fatalf("failed to create output file: %v", err)
	}

	err = about.About().Render(context.Background(), f)
	if err != nil {
		log.Fatalf("failed to write output file: %v", err)
	}
}
