package main

import (
	"bytes"
	"fmt"
	"log"
	"os"

	"github.com/yuin/goldmark"
)

func MarkdownToHTML(markdown []byte) []byte {
	var buf bytes.Buffer
	if err := goldmark.Convert([]byte(markdown), &buf); err != nil {
		log.Fatal(err)
	}

	return buf.Bytes()
}

func main() {
	inputPath := "io/01-input.md"
	outputPath := "io/output.html"

	mdBytes, err := os.ReadFile(inputPath)
	if err != nil {
		log.Fatalf("Failed to read %s: %v", inputPath, err)
	}

	// Convert Markdown to HTML
	html := MarkdownToHTML(mdBytes)
	if err := os.WriteFile(outputPath, html, 0644); err != nil {
		log.Fatalf("Failed to write %s: %v", outputPath, err)
	}

	fmt.Println("HTML written to ", outputPath)
}
