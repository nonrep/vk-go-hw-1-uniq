package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type Test struct {
	input    []string
	options  Options
	expected []string
}

var testsOK = []Test{
	{
		input:    defaultInput,
		options:  Options{},
		expected: defaultResult,
	},
	{
		input:    countInput,
		options:  Options{Count: true},
		expected: countResult,
	},
	{
		input:    duplicateInput,
		options:  Options{Duplicate: true},
		expected: duplicateResult,
	},
	{
		input:    uniqueInput,
		options:  Options{Unique: true},
		expected: uniqueResult,
	},
	{
		input:    ignoreCaseInput,
		options:  Options{IgnoreCase: true},
		expected: ignoreCaseResult,
	},
	{
		input:    skipFieldsInput,
		options:  Options{SkipFields: 1},
		expected: skipFieldsResult,
	},
	{
		input:    skipCharsInput,
		options:  Options{SkipChars: 1},
		expected: skipCharsResult,
	},
	{
		input:    skipFieldsCharsInput,
		options:  Options{SkipFields: 1, SkipChars: 1},
		expected: skipFieldsCharsResult,
	},
}

func TestOK(t *testing.T) {
	for _, test := range testsOK {
		t.Run("", func(t *testing.T) {
			result, err := uniq(test.input, test.options)
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
		options: Options{Count: true, Duplicate: true},
	},
	{
		input:   defaultInput,
		options: Options{Count: true, Unique: true},
	},
	{
		input:   defaultInput,
		options: Options{Unique: true, Duplicate: true},
	},
	{
		input:   defaultInput,
		options: Options{SkipFields: -10},
	},
	{
		input:   defaultInput,
		options: Options{SkipChars: -1},
	},
}

func TestFail(t *testing.T) {
	for _, test := range testsFail {
		t.Run("", func(t *testing.T) {
			_, err := uniq(test.input, test.options)
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

var countInput = []string{
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

var countResult = []string{
	"3 I love music.",
	"1 ",
	"2 I love music of Kartik.",
	"1 Thanks.",
	"2 I love music of Kartik.",
}

var duplicateInput = []string{
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

var duplicateResult = []string{
	"I love music.",
	"I love music of Kartik.",
	"I love music of Kartik.",
}

var uniqueInput = []string{
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

var uniqueResult = []string{
	"",
	"Thanks.",
}

var ignoreCaseInput = []string{
	"I LOVE MUSIC.",
	"I love music.",
	"I LoVe MuSiC.",
	"",
	"I love MuSIC of Kartik.",
	"I love music of kartik.",
	"Thanks.",
	"I love music of kartik.",
	"I love MuSIC of Kartik.",
}

var ignoreCaseResult = []string{
	"I LOVE MUSIC.",
	"",
	"I love MuSIC of Kartik.",
	"Thanks.",
	"I love music of kartik.",
}

var skipFieldsInput = []string{
	"We love music.",
	"I love music.",
	"They love music.",
	"",
	"I love music of Kartik.",
	"We love music of Kartik.",
	"Thanks.",
}

var skipFieldsResult = []string{
	"We love music.",
	"",
	"I love music of Kartik.",
	"Thanks.",
}

var skipCharsInput = []string{
	"I love music.",
	"A love music.",
	"C love music.",
	"",
	"I love music of Kartik.",
	"We love music of Kartik.",
	"Thanks.",
}

var skipCharsResult = []string{
	"I love music.",
	"",
	"I love music of Kartik.",
	"We love music of Kartik.",
	"Thanks.",
}

var skipFieldsCharsInput = []string{
	"One I love music.",
	"Two A love music.",
	"Three C love music.",
	"",
	"I love music of Kartik.",
	"We love music of Kartik.",
	"Thanks.",
}

var skipFieldsCharsResult = []string{
	"One I love music.",
	"",
	"I love music of Kartik.",
	"Thanks.",
}
