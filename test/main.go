package main

import (
	"bytes"
	"fmt"
        "os"

	"github.com/oss-experiment-uec/2024-h1910541-pdf"
)

func main() {
	// pdf.DebugOn = true //
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <pdf-path>")
		os.Exit(1)
	}

	pdfPath := os.Args[1]

	content, err := readPdf(pdfPath)
	if err != nil {
		fmt.Printf("Error reading PDF: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(content)
        wordCount := pdf.CountWords(content)
        fmt.Printf("The PDF contains %d words.\n", wordCount)
        return

}


func readPdf(path string) (string, error) {
	f, r, err := pdf.Open(path)
	// remember close file
    defer f.Close()
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
    b, err := r.GetPlainText()
    if err != nil {
        return "", err
    }
    buf.ReadFrom(b)
	return buf.String(), nil
}
