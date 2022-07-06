package main

import (
	"fmt"
	"os"
	"unicode"
)

func main() {
	count := 0

	for i := 1; i < len(os.Args); i++ {
		count += wordCount(os.Args[i])
	}

	fmt.Println(count)
}

func wordCount(text string) int {
	count := 0
	prev := false

	for _, ch := range text {
		cur := isWord(ch)

		if !prev && cur {
			count++
		}

		prev = cur
	}

	return count
}

func isWord(ch rune) bool {
	// Разделителем слов считаем только пробел.
	// Соответственно, слово — любая последовательность символов, не включающая пробелов.
	return !unicode.IsSpace(ch)
}
