package uniq_test

import (
	"testing"

	"github.com/nonrep/go-homework-1-uniq/uniq"
	"github.com/stretchr/testify/require"
)

type Test struct {
	input    []string
	options  uniq.Options
	expected []string
}

var testsOK = []Test{
	{
		input:    defaultInput,
		options:  uniq.Options{},
		expected: defaultResult,
	},
	{
		input: []string{
			"I love music.",
			"I love music.",
			"I love music.",
			"",
			"I love music of Kartik.",
			"I love music of Kartik.",
			"Thanks.",
			"I love music of Kartik.",
			"I love music of Kartik.",
		},
		options: uniq.Options{Count: true},
		expected: []string{
			"3 I love music.",
			"1 ",
			"2 I love music of Kartik.",
			"1 Thanks.",
			"2 I love music of Kartik.",
		},
	},
	{
		input: []string{
			"I love music.",
			"I love music.",
			"I love music.",
			"",
			"I love music of Kartik.",
			"I love music of Kartik.",
			"Thanks.",
			"I love music of Kartik.",
			"I love music of Kartik.",
		},
		options: uniq.Options{Duplicate: true},
		expected: []string{
			"I love music.",
			"I love music of Kartik.",
			"I love music of Kartik.",
		},
	},
	{
		input: []string{
			"I love music.",
			"I love music.",
			"I love music.",
			"",
			"I love music of Kartik.",
			"I love music of Kartik.",
			"Thanks.",
			"I love music of Kartik.",
			"I love music of Kartik.",
		},
		options: uniq.Options{Unique: true},
		expected: []string{
			"",
			"Thanks.",
		},
	},
	{
		input: []string{
			"I LOVE MUSIC.",
			"I love music.",
			"I LoVe MuSiC.",
			"",
			"I love MuSIC of Kartik.",
			"I love music of kartik.",
			"Thanks.",
			"I love music of kartik.",
			"I love MuSIC of Kartik.",
		},
		options: uniq.Options{IgnoreCase: true},
		expected: []string{
			"I LOVE MUSIC.",
			"",
			"I love MuSIC of Kartik.",
			"Thanks.",
			"I love music of kartik.",
		},
	},
	{
		input: []string{
			"We love music.",
			"I love music.",
			"They love music.",
			"",
			"I love music of Kartik.",
			"We love music of Kartik.",
			"Thanks.",
		},
		options: uniq.Options{SkipFields: 1},
		expected: []string{
			"We love music.",
			"",
			"I love music of Kartik.",
			"Thanks.",
		},
	},
	{
		input: []string{
			"I love music.",
			"A love music.",
			"C love music.",
			"",
			"I love music of Kartik.",
			"We love music of Kartik.",
			"Thanks.",
		},
		options: uniq.Options{SkipChars: 1},
		expected: []string{
			"I love music.",
			"",
			"I love music of Kartik.",
			"We love music of Kartik.",
			"Thanks.",
		},
	},
	{
		input: []string{
			"One I love music.",
			"Two A love music.",
			"Three C love music.",
			"",
			"I love music of Kartik.",
			"We love music of Kartik.",
			"Thanks.",
		},
		options: uniq.Options{SkipFields: 1, SkipChars: 1},
		expected: []string{
			"One I love music.",
			"",
			"I love music of Kartik.",
			"Thanks.",
		},
	},
}

func TestOK(t *testing.T) {
	for _, test := range testsOK {
		t.Run("", func(t *testing.T) {
			result, err := uniq.Uniq(test.input, test.options)
			if err != nil {
				t.Errorf("Error: %s", err)
				return
			}
			require.Equal(t, result, test.expected, "The two words should be the same.")
		})
	}
}

var testsFail = []Test{
	{
		input:   defaultInput,
		options: uniq.Options{Count: true, Duplicate: true},
	},
	{
		input:   defaultInput,
		options: uniq.Options{Count: true, Unique: true},
	},
	{
		input:   defaultInput,
		options: uniq.Options{Unique: true, Duplicate: true},
	},
	{
		input:   defaultInput,
		options: uniq.Options{SkipFields: -10},
	},
	{
		input:   defaultInput,
		options: uniq.Options{SkipChars: -1},
	},
}

func TestFail(t *testing.T) {
	for _, test := range testsFail {
		t.Run("", func(t *testing.T) {
			_, err := uniq.Uniq(test.input, test.options)
			require.Error(t, err, "Expected an error, but got none.")
		})
	}
}

var defaultInput = []string{
	"I love music.",
	"I love music.",
	"I love music.",
	"",
	"I love music of Kartik.",
	"I love music of Kartik.",
	"Thanks.",
	"I love music of Kartik.",
	"I love music of Kartik.",
}

var defaultResult = []string{
	"I love music.",
	"",
	"I love music of Kartik.",
	"Thanks.",
	"I love music of Kartik.",
}
