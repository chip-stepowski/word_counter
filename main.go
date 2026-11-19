package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {

	text := `orem ipsum dolor sit amet, consectetur adipiscing elit. Nulla hendrerit eros eu ipsum blandit condimentum. Proin at varius lectus. In tristique purus a odio posuere, eget volutpat urna interdum. Curabitur non leo urna. Nam augue tortor, pulvinar quis ligula in, lacinia lobortis odio. Donec vel sodales sapien, ut vehicula sem. Vestibulum lacinia mattis libero, ut eleifend nibh mollis et. Aliquam vitae eros arcu.`

	// lowercase
	lower := strings.ToLower(text)

	// remove punctuation
	splitFunc := func(r rune) bool {
		return unicode.IsSpace(r) || r == ',' || r == '.' || r == ';'
	}
	// returns an array
	rpunc := strings.FieldsFunc(lower, splitFunc)
	fmt.Println(rpunc)

}

// - Hardcode a paragraph of text in your program.
// - Lowercase it, split it into words (`strings.Fields`), and strip punctuation (`strings.Trim`).
// - Count each word’s frequency in a `map[string]int`.
// - Print every word with its count.

// **Done when:** counts are correct for a paragraph you verified by hand.
