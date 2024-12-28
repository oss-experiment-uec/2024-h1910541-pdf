package main

import (
	"bytes"
	"fmt"

	"github.com/ledongthuc/pdf"
)

func main() {
	// pdf.DebugOn = true //
	content, err := readPdf("sample2.pdf") // Read local pdf file
	if err != nil {
		panic(err)
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
