package toc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSlugify(t *testing.T) {
	testCases := map[string]struct {
		in       string
		expected string
	}{
		"applies lower case":    {in: "MysTrInghEre", expected: "mystringhere"},
		"replace space with -":  {in: "Some ex ample", expected: "some-ex-ample"},
		"replace () with space": {in: "Header (something)", expected: "header-something"},
		"replace [] with space": {in: "Header [something]", expected: "header-something"},
		"replace {} with space": {in: "Header {something}", expected: "header-something"},
		"replace \" with space": {in: "Header \"something\"", expected: "header-something"},
		"replace ' with space":  {in: "Header 'something'", expected: "header-something"},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, testCase.expected, slugify(testCase.in))
		})
	}
}
