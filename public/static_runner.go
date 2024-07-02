package main

import (
    "github.com/harish876/portfolio/views/home"
    "context"
    "log"
    "os"
    "github.com/harish876/portfolio/views/about"
)

//RUN THIS FILE TO REGENERATE THE STATIC PAGES

func main() {
    f_0, err := os.Create("public/about.html")
    if err != nil {
        log.Fatalf("failed to create output file 'public/about.html': %v", err)
    }
    defer f_0.Close()

    err = about.About().Render(context.Background(), f_0)
    if err != nil {
        log.Fatalf("failed to write output file 'public/about.html': %v", err)
    }
    log.Printf("public/about.html created successfully.")
    f_1, err := os.Create("public/home.html")
    if err != nil {
        log.Fatalf("failed to create output file 'public/home.html': %v", err)
    }
    defer f_1.Close()

    err = home.Home().Render(context.Background(), f_1)
    if err != nil {
        log.Fatalf("failed to write output file 'public/home.html': %v", err)
    }
    log.Printf("public/home.html created successfully.")
}
