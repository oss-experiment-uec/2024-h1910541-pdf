package pdf

import (
	"regexp"
)

// CountWords counts the number of words in a given string
func CountWords(text string) int {
	re := regexp.MustCompile(`[^\s]+`)
	words := re.FindAllString(text, -1)
	return len(words)
}
