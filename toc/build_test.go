package toc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuild(t *testing.T) {
	testCases := map[string]struct {
		data        []byte
		header      string
		depth       int
		skipHeaders int
		addHeader   bool
		expectedToC []string
	}{
		"success - empty markdown": {
			data:        []byte(``),
			header:      "# Table of Contents",
			depth:       0,
			skipHeaders: 0,
			addHeader:   true,
			expectedToC: []string{
				"<!-- ToC start -->",
				"# Table of Contents\n",
				"<!-- ToC end -->",
			},
		},
		"success - full example": {
			data: []byte(`
Header 1
========

Some content

Header 2
--------

Some content

# Header 3

Some content

## Header 4

Some content

### Header 5

Some content

#### Header 6

Some content

#### Header 7

Some content
`),
			header:      "# Table of Contents",
			depth:       0,
			skipHeaders: 0,
			addHeader:   true,
			expectedToC: []string{
				"<!-- ToC start -->",
				"# Table of Contents\n",
				"1. [Header 1](#header-1)",
				"   1. [Header 2](#header-2)",
				"1. [Header 3](#header-3)",
				"   1. [Header 4](#header-4)",
				"      1. [Header 5](#header-5)",
				"         1. [Header 6](#header-6)",
				"         1. [Header 7](#header-7)",
				"<!-- ToC end -->",
			},
		},
		"success - Repeated headers": {
			data: []byte(`
Header 1
========
Content

# Header 1
Some content
`),
			header:      "# Table of Contents",
			depth:       0,
			skipHeaders: 0,
			addHeader:   true,
			expectedToC: []string{
				"<!-- ToC start -->",
				"# Table of Contents\n",
				"1. [Header 1](#header-1)",
				"1. [Header 1](#header-1-1)",
				"<!-- ToC end -->",
			},
		},
		"skipping 3 headers": {
			data: []byte(`
Header 1
========

Some content

Header 2
--------

Some content

# Header 3

Some content

## Header 4

Some content

### Header 5

Some content

#### Header 6

Some content
`),
			header:      "# Table of Contents",
			depth:       0,
			skipHeaders: 3,
			addHeader:   true,
			expectedToC: []string{
				"<!-- ToC start -->",
				"# Table of Contents\n",
				"   1. [Header 4](#header-4)",
				"      1. [Header 5](#header-5)",
				"         1. [Header 6](#header-6)",
				"<!-- ToC end -->",
			},
		},
		"only render first level headers": {
			data: []byte(`
Header 1
========

Some content

Header 2
--------

Some content

# Header 3

Some content

## Header 4

Some content

### Header 5

Some content

#### Header 6

Some content
`),
			header:      "# Table of Contents",
			depth:       1,
			skipHeaders: 0,
			addHeader:   true,
			expectedToC: []string{
				"<!-- ToC start -->",
				"# Table of Contents\n",
				"1. [Header 1](#header-1)",
				"1. [Header 3](#header-3)",
				"<!-- ToC end -->",
			},
		},
		"only render 2 levels of headers": {
			data: []byte(`
Header 1
========

Some content

Header 2
--------

Some content

# Header 3

Some content

## Header 4

Some content

### Header 5

Some content

#### Header 6

Some content
`),
			header:      "# Table of Contents",
			depth:       2,
			skipHeaders: 0,
			addHeader:   true,
			expectedToC: []string{
				"<!-- ToC start -->",
				"# Table of Contents\n",
				"1. [Header 1](#header-1)",
				"   1. [Header 2](#header-2)",
				"1. [Header 3](#header-3)",
				"   1. [Header 4](#header-4)",
				"<!-- ToC end -->",
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			toc, err := Build(
				testCase.data,
				testCase.header,
				testCase.depth,
				testCase.skipHeaders,
				testCase.addHeader,
			)
			assert.Nil(t, err)
			assert.Equal(t, testCase.expectedToC, toc)
		})
	}
}
