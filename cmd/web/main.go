package main

import (
	"bytes"
	"fmt"
	"log"
	"os"

	"github.com/yuin/goldmark"
)

func MarkdownToHTML(markdown string) string {
	var buf bytes.Buffer
	if err := goldmark.Convert([]byte(markdown), &buf); err != nil {
		log.Fatal(err)
	}

	return buf.String()
}

func main() {
	// Markdown sample
	markdown := "# Hello, World!\nThis is a **Markdown** test."

	// Convert Markdown to HTML
	html := MarkdownToHTML(markdown)

	// Print result
	fmt.Println(html)

	err := os.WriteFile("output.html", []byte(html), 0644)
	if err != nil {
		log.Fatal(err)
	}
}
