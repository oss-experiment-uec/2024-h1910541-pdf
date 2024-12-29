package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	
	rand.Seed(time.Now().UnixNano())

	wordCount := rand.Intn(1000) + 1

	words := make([]string, wordCount)
	for i := 0; i < wordCount; i++ {
		words[i] = randomWord()
	}

	file, err := os.Create("sample.txt")
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		return
	}
	defer file.Close()

	
	file.WriteString(strings.Join(words, " "))

	fmt.Printf("Generated sample.txt with %d words.\n", wordCount)
}

func randomWord() string {
	letters := "abcdefghijklmnopqrstuvwxyz"
	wordLength := rand.Intn(10) + 1
	word := make([]byte, wordLength)
	for i := 0; i < wordLength; i++ {
		word[i] = letters[rand.Intn(len(letters))]
	}
	return string(word)
}
