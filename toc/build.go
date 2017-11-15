package toc

import (
	"bufio"
	"bytes"
	"fmt"
	"regexp"
	"strings"
)

var (
	rHashHeader        = regexp.MustCompile("^(?P<indent>#+) ?(?P<title>.+)$")
	rUnderscoreHeader1 = regexp.MustCompile("^=+$")
	rUnderscoreHeader2 = regexp.MustCompile("^\\-+$")
)

// Build is returning a ToC based on the input markdown.
func Build(d []byte, header string, depth, skipHeaders int, addHeader bool) ([]string, error) {
	toc := []string{
		"<!-- ToC start -->",
	}

	if addHeader {
		toc = append(toc, fmt.Sprintf("%s\n", header))
	}

	var previousLine string
	s := bufio.NewScanner(bytes.NewReader(d))
	for s.Scan() {
		switch {
		case rHashHeader.Match(s.Bytes()):
			m := rHashHeader.FindStringSubmatch(s.Text())
			if depth > 0 && len(m[1]) > depth {
				continue
			}

			indent := len(m[1]) - 1
			title := m[2]

			if skipHeaders > 0 {
				skipHeaders--
				continue
			}

			toc = append(toc, fmt.Sprintf("%s- [%s](#%s)", strings.Repeat("  ", indent), title, slugify(title)))

		case rUnderscoreHeader1.Match(s.Bytes()):
			if skipHeaders > 0 {
				skipHeaders--
				continue
			}

			toc = append(toc, fmt.Sprintf("- [%s](#%s)", previousLine, slugify(previousLine)))

		case rUnderscoreHeader2.Match(s.Bytes()):
			if depth > 0 && depth < 2 {
				continue
			}

			if skipHeaders > 0 {
				skipHeaders--
				continue
			}

			toc = append(toc, fmt.Sprintf("  - [%s](#%s)", previousLine, slugify(previousLine)))
		}

		previousLine = s.Text()
	}
	if err := s.Err(); err != nil {
		return []string{}, err
	}

	toc = append(toc, "<!-- ToC end -->")

	return toc, nil
}
