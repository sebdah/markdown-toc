package toc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReplace(t *testing.T) {
	testCases := map[string]struct {
		data     []byte
		toc      []string
		expected []string
	}{
		"without existing ToC": {
			data: []byte(`# Header 1
Some content

## Header 2
More content
`),
			toc: []string{
				"<!-- ToC start -->",
				"# Table of Contents\n",
				"1. [Header 1](#header-1)",
				"   1. [Header 2](#header-2)",
				"<!-- ToC end -->\n",
			},
			expected: []string{
				"<!-- ToC start -->",
				"# Table of Contents\n",
				"1. [Header 1](#header-1)",
				"   1. [Header 2](#header-2)",
				"<!-- ToC end -->\n",
				"# Header 1",
				"Some content",
				"",
				"## Header 2",
				"More content",
			},
		},
		"with existing ToC": {
			data: []byte(`# Header 1
Some content

<!-- ToC start -->
data
<!-- ToC end -->

## Header 2
More content
`),
			toc: []string{
				"<!-- ToC start -->",
				"# Table of Contents\n",
				"1. [Header 1](#header-1)",
				"   1. [Header 2](#header-2)",
				"<!-- ToC end -->\n",
			},
			expected: []string{
				"# Header 1",
				"Some content",
				"",
				"<!-- ToC start -->",
				"# Table of Contents\n",
				"1. [Header 1](#header-1)",
				"   1. [Header 2](#header-2)",
				"<!-- ToC end -->\n",
				"",
				"## Header 2",
				"More content",
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, testCase.expected, Replace(testCase.data, testCase.toc))
		})
	}
}
