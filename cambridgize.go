// Cambridgize shuffles the letters of all words in a text except first and last.
//
// Usage:
//  go run cambridgize.go "My text to be cambridgized"
//
// Inspired by this famous article:
//  "Aoccdrnig to a rscheearch at Cmabrigde Uinervtisy, it deosn't mttaer
//  in waht oredr the ltteers in a wrod are, the olny iprmoetnt tihng is
//  taht the frist and lsat ltteer be at the rghit pclae. [...]"
package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"time"
)

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
		return
	}
}

func run() error {
	if len(os.Args) < 2 {
		return errors.New("missing text string")
	}
	text := os.Args[1]
	fmt.Println(Cambridgize(text))

	return nil
}

// Cambridgize shuffles the inner letters of every word in a given text
// and returns a new text.
func Cambridgize(text string) string {
	rgx := wordRegexp()
	return rgx.ReplaceAllStringFunc(text, cambridgizeWord)
}

// cambridgzizeWord shuffles the inner letters of a word,
// e.g. all letters except the first and the last ones.
func cambridgizeWord(word string) string {
	randRange := func(a, b int) int {
		return a + rand.New(rand.NewSource(time.Now().UnixNano())).Intn(b-a)
	}

	runes := []rune(word)
	n := len(runes)
	min, max := 1, n-1

	// First and last letters stay in place,
	// so words <= 3 letters are unaltered.
	if n <= 3 {
		return word
	}

	for c := min; c < max; c++ {
		curr := &runes[c]
		swap := &runes[randRange(min, max)]
		*curr, *swap = *swap, *curr
	}

	return string(runes)
}

// wordRegexp returns a regexp that identifies words in a string.
func wordRegexp() *regexp.Regexp {
	// Can't rely on '\w' to define a word character since it does not include
	// accents. Instead, define a word character by being:
	// - not a digit (let's not shuffle numbers!)
	// - not a separator, which include: space-like, apostrophe, dash.
	// - not an element of punctuation: ? ! . , ; :
	//
	// Words <= 3 letters are also excluded, since first and last letter
	// should not move.
	pattern := `[^0-9\s'-\?!\.,;:]{3,}`
	return regexp.MustCompile(pattern)
}
