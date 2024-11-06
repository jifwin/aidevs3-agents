package main

import (
    "fmt"
    "log"
    "net/http"
    "os"

    "github.com/PuerkitoBio/goquery"
)

func main() {
    // Check if URL is provided as an argument
    if len(os.Args) != 2 {
        fmt.Println("Usage: go run main.go <URL>")
        os.Exit(1)
    }

    url := os.Args[1]

    // Send a GET request to the URL
    resp, err := http.Get(url)
    if err != nil {
        log.Fatalf("Error fetching URL: %v", err)
    }
    defer resp.Body.Close()

    // Check if the HTTP status is OK
    if resp.StatusCode != http.StatusOK {
        log.Fatalf("Error: received status code %d", resp.StatusCode)
    }

    // Load the HTML document
    doc, err := goquery.NewDocumentFromReader(resp.Body)
    if err != nil {
        log.Fatalf("Error parsing HTML: %v", err)
    }

    // Find the paragraph with id="human-question"
    p := doc.Find("p#human-question")
    if p.Length() == 0 {
        fmt.Println("No paragraph found with id 'human-question'")
        os.Exit(1)
    }

    // Extract and print the text content of the paragraph
    text := p.Text()
    fmt.Println("Extracted paragraph content:")
    fmt.Println(text)
}
