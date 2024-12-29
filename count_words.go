package pdf

import (
	"regexp"
)

// CountWords counts the number of words in a given string
func CountWords(text string) int {

<<<<<<< HEAD
	re := regexp.MustCompile(`[^\s]+`)
=======
	re := regexp.MustCompile(`\b\w+\b`)
>>>>>>> e14142b8abd99ef3231544b9d1337ef83ba8749d
	words := re.FindAllString(text, -1)
	return len(words)
}
