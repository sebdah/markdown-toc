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
	rCodeblocks        = regexp.MustCompile("^```")
)

// Build is returning a ToC based on the input markdown.
func Build(d []byte, header string, depth, skipHeaders int, addHeader bool) ([]string, error) {
	toc := []string{
		"<!-- ToC start -->",
	}

	if addHeader {
		toc = append(toc, fmt.Sprintf("%s\n", header))
	}

	seenHeaders := make(map[string]int)
	var previousLine string

	appendToC := func(title string, indent int) {
		link := slugify(title)

		if skipHeaders > 0 {
			skipHeaders--
			return
		}

		if _, ok := seenHeaders[link]; ok {
			seenHeaders[link]++
			link = fmt.Sprintf("%s-%d", link, seenHeaders[link]-1)
		} else {
			seenHeaders[link] = 1
		}

		toc = append(toc, fmt.Sprintf("%s1. [%s](#%s)", strings.Repeat("   ", indent), title, link))
	}

	s := bufio.NewScanner(bytes.NewReader(d))
	for s.Scan() {
		switch {
		case rHashHeader.Match(s.Bytes()):
			m := rHashHeader.FindStringSubmatch(s.Text())
			if depth > 0 && len(m[1]) > depth {
				continue
			}

			appendToC(m[2], len(m[1])-1)

		case rUnderscoreHeader1.Match(s.Bytes()):
			appendToC(previousLine, 0)

		case rUnderscoreHeader2.Match(s.Bytes()):
			if depth > 0 && depth < 2 {
				continue
			}

			appendToC(previousLine, 1)
		case rCodeblocks.Match(s.Bytes()):
			toc = append(toc, s.Text())
		}

		previousLine = s.Text()
	}
	if err := s.Err(); err != nil {
		return []string{}, err
	}

	toc = append(toc, "<!-- ToC end -->")

	toc = dropCodeblocks(toc)

	return toc, nil
}

func dropCodeblocks(toc []string) []string {

	var result []string

	var codeStart, codeEnd int

	for i, l := 0, len(toc); i < l; i++ {

		if codeEnd > 0 {
			i = codeEnd
			codeEnd = 0
			continue
		}

		if rCodeblocks.MatchString(toc[i]) {
			codeStart = i
			toc[i] = ""

			for j := codeStart; j < l; j++ {

				if codeEnd > 0 {
					break
				}

				if rCodeblocks.MatchString(toc[j]) {
					codeEnd = j
					toc[j] = ""

					for k := codeStart; k < codeEnd; k++ {
						toc[k] = ""
					}
				}

			}
		}

	}

	for _, v := range toc {

		if v != "" {
			result = append(result, v)
		}

	}

	return result

}
