package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// strings are read-only slices of bytes, encoded in UTF-8
	// char -> rune (represents a Unicode codepoint)

	const s = "สวัสดี"

	// amount of raw bytes
	fmt.Println("len:", len(s))
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()

	fmt.Println("rune count:", utf8.RuneCountInString(s))

	// range decodes runes
	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}

	fmt.Println("\nUsing DecodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width

		examineRune(runeValue)
	}
}

func examineRune(r rune) {
	if r == 't' {
		fmt.Println("found tee")
	} else if r == 'ส' {
		fmt.Println("found so sua")
	}
}
