// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
	"fmt"
	"log"
	"regexp"
	"strings"
	"unicode"
)

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	abbr := ""
	reg, err := regexp.Compile("-")
	if err != nil {
		log.Fatal(err)
	}

	processedString := reg.ReplaceAllString(s, " ")
	fmt.Println(processedString)

	splitedwords := strings.Split(processedString, " ")
	fmt.Println(splitedwords)

	for _, w := range splitedwords {

		for _, t := range w {
			if unicode.IsLetter(t) {
				abbr += string(t)
				break
			}
		}
	}
	abbr = strings.ToUpper(abbr)
	return abbr
}
