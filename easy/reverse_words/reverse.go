package reverse_words

import "strings"

func reverseWords(s string) string {
	words := strings.Fields(s)
	b := strings.Builder{}

	b.Grow(len(words))

	for i := len(words) - 1; i >= 0; i-- {
		if i != 0 {
			b.WriteString(words[i])
			b.WriteString(" ")
			continue
		}

		b.WriteString(words[i])

	}

	return b.String()
}
