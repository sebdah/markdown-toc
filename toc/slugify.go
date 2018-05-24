package toc

import "strings"
import "unicode"

// slugify is converting a string into a slug representation of the string. The
// rules here are adapted to how GitHub is creating slugs from the headers.
func slugify(s string) string {
	droppedChars := []string{
		"\"", "'", "`", ".",
		"!", ",", "~", "&",
		"%", "^", "*", "#",
		"@", "|",
		"(", ")",
		"{", "}",
		"[", "]",
	}

	s = strings.ToLower(s)

	for _, c := range droppedChars {
		s = strings.Replace(s, c, "", -1)
	}
	f := func(r rune) bool {
		return unicode.IsSpace(r)
	}
	s = strings.TrimRightFunc(s, f)

	s = strings.Replace(s, " ", "-", -1)

	return s
}
