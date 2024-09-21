package main

import (
	"bufio"
	"bytes"
	"os"
	"testing"
)

var defaultInput = `I love music.
I love music.
I love music.

I love music of Kartik.
I love music of Kartik.
Thanks.
I love music of Kartik.
I love music of Kartik.
`

var defaultResult = `I love music.

I love music of Kartik.
Thanks.
I love music of Kartik.
`

var countInput = `I love music.
I love music.
I love music.

I love music of Kartik.
I love music of Kartik.
Thanks.
I love music of Kartik.
I love music of Kartik.
`

var countResult = `3 I love music.
1 
2 I love music of Kartik.
1 Thanks.
2 I love music of Kartik.
`

var duplicateInput = `I love music.
I love music.
I love music.

I love music of Kartik.
I love music of Kartik.
Thanks.
I love music of Kartik.
I love music of Kartik.
`

var duplicateResult = `I love music.
I love music of Kartik.
I love music of Kartik.
`

var uniqueInput = `I love music.
I love music.
I love music.

I love music of Kartik.
I love music of Kartik.
Thanks.
I love music of Kartik.
I love music of Kartik.
`

var uniqueResult = `
Thanks.
`

var ignoreCaseInput = `I LOVE MUSIC.
I love music.
I LoVe MuSiC.

I love MuSIC of Kartik.
I love music of kartik.
Thanks.
I love music of kartik.
I love MuSIC of Kartik.
`

var ignoreCaseResult = `I LOVE MUSIC.

I love MuSIC of Kartik.
Thanks.
I love music of kartik.
`
var skipFieldsInput = `We love music.
I love music.
They love music.

I love music of Kartik.
We love music of Kartik.
Thanks.
`

var skipFieldsResult = `We love music.

I love music of Kartik.
Thanks.
`

var skipCharsInpit = `I love music.
A love music.
C love music.

I love music of Kartik.
We love music of Kartik.
Thanks.
`

var skipCharsResult = `I love music.

I love music of Kartik.
We love music of Kartik.
Thanks.
`

var skipFieldsCharsInput = `One I love music.
Two A love music.
Three C love music.

I love music of Kartik.
We love music of Kartik.
Thanks.
`

var skipFieldsCharsResult = `One I love music.

I love music of Kartik.
Thanks.
`

func TestOkDefault(t *testing.T) {
	in := bytes.NewBufferString(defaultInput)
	out := bytes.NewBuffer(nil)
	options := Options{}
	err := uniq(in, out, options)
	if err != nil {
		t.Errorf("Testing default failed: %s", err)
	}
	result := out.String()
	if result != defaultResult {
		t.Error("Testing default failed, result not match")
	}
}

func TestOkInputFile(t *testing.T) {
	in, err := os.Open("testdata/input.txt")
	if err != nil {
		t.Errorf("Testing input file failed: %s", err)
	}
	defer in.Close()

	out := bytes.NewBuffer(nil)

	options := Options{}

	err = uniq(in, out, options)
	if err != nil {
		t.Errorf("Testing input file failed: %s", err)
	}

	result := out.String()
	if result != defaultResult {
		t.Error("Testing input file failed, result not match")
	}
}

func TestFailInputFile(t *testing.T) {
	in, err := os.Open("NotExistsFile.txt")
	if err == nil {
		t.Errorf("Testing wrong input file failed: expected error")
	}
	defer in.Close()
}

func TestOkInputOutputFiles(t *testing.T) {
	in, err := os.Open("testdata/input.txt")
	if err != nil {
		t.Errorf("Testing input/output files failed: %s", err)
	}
	defer in.Close()

	out, err := os.Create("testdata/output.txt")
	if err != nil {
		t.Errorf("Testing input/output files failed: %s", err)
	}

	options := Options{}

	err = uniq(in, out, options)
	if err != nil {
		t.Errorf("Testing input/output files failed: %s", err)
	}

	var result string
	out.Close()
	out, err = os.Open("testdata/output.txt")
	if err != nil {
		t.Errorf("Testing input/output files failed: %s", err)
	}
	scanner := bufio.NewScanner(out)
	for scanner.Scan() {
		result += scanner.Text() + "\n"
	}

	if result != defaultResult {
		t.Error("Testing input/output files failed, result not match")
	}
}

func TestOkCount(t *testing.T) {
	in := bytes.NewBufferString(countInput)
	out := bytes.NewBuffer(nil)
	options := Options{Count: true}
	err := uniq(in, out, options)
	if err != nil {
		t.Errorf("Testing count failed: %s", err)
	}
	result := out.String()
	if result != countResult {
		t.Error("Testing count failed, result not match")
	}
}

func TestOkDuplicate(t *testing.T) {
	in := bytes.NewBufferString(duplicateInput)
	out := bytes.NewBuffer(nil)
	options := Options{Duplicate: true}
	err := uniq(in, out, options)
	if err != nil {
		t.Errorf("Testing duplicate failed: %s", err)
	}
	result := out.String()
	if result != duplicateResult {
		t.Error("Testing duplicate failed, result not match")
	}
}

func TestOkUnique(t *testing.T) {
	in := bytes.NewBufferString(uniqueInput)
	out := bytes.NewBuffer(nil)
	options := Options{Unique: true}
	err := uniq(in, out, options)
	if err != nil {
		t.Errorf("Testing unique failed: %s", err)
	}
	result := out.String()
	if result != uniqueResult {
		t.Error("Testing unique failed, result not match")
	}
}

func TestOkIgnoreCase(t *testing.T) {
	in := bytes.NewBufferString(ignoreCaseInput)
	out := bytes.NewBuffer(nil)
	options := Options{IgnoreCase: true}
	err := uniq(in, out, options)
	if err != nil {
		t.Errorf("Testing ignore case failed: %s", err)
	}
	result := out.String()
	if result != ignoreCaseResult {
		t.Error("Testing ignore case failed, result not match")
	}
}

func TestOkSkipFields(t *testing.T) {
	in := bytes.NewBufferString(skipFieldsInput)
	out := bytes.NewBuffer(nil)
	options := Options{SkipFields: 1}
	err := uniq(in, out, options)
	if err != nil {
		t.Errorf("Testing skip fields failed: %s", err)
	}
	result := out.String()
	if result != skipFieldsResult {
		t.Error("Testing skip fields failed, result not match")
	}
}

func TestOkSkipChars(t *testing.T) {
	in := bytes.NewBufferString(skipCharsInpit)
	out := bytes.NewBuffer(nil)
	options := Options{SkipChars: 1}
	err := uniq(in, out, options)
	if err != nil {
		t.Errorf("Testing skip chars failed: %s", err)
	}
	result := out.String()
	if result != skipCharsResult {
		t.Error("Testing skip chars failed, result not match")
	}
}

func TestOkSkipFieldsChars(t *testing.T) {
	in := bytes.NewBufferString(skipFieldsCharsInput)
	out := bytes.NewBuffer(nil)
	options := Options{SkipFields: 1, SkipChars: 1}
	err := uniq(in, out, options)
	if err != nil {
		t.Errorf("Testing skip fields and chars failed: %s", err)
	}
	result := out.String()
	if result != skipFieldsCharsResult {
		t.Error("Testing skip fields and chars failed, result not match")
	}
}

func TestFailWrong(t *testing.T) {
	in, err := os.Open("NotExistsFile.txt")
	if err == nil {
		t.Errorf("Testing wrong input file failed: expected error")
	}
	defer in.Close()
}

func TestFailWrongOptions(t *testing.T) {
	in := bytes.NewBufferString(defaultInput)
	out := bytes.NewBuffer(nil)
	options := Options{Count: true, Duplicate: true}
	err := uniq(in, out, options)
	if err == nil {
		t.Error("Testing wrong options failed: expected error")
	}
	options = Options{Count: true, Unique: true}
	err = uniq(in, out, options)
	if err == nil {
		t.Error("Testing wrong options failed: expected error")
	}
	options = Options{Unique: true, Duplicate: true}
	err = uniq(in, out, options)
	if err == nil {
		t.Error("Testing wrong options failed: expected error")
	}
	options = Options{SkipFields: -10}
	err = uniq(in, out, options)
	if err == nil {
		t.Error("Testing wrong options failed: expected error")
	}
	options = Options{SkipChars: -1}
	err = uniq(in, out, options)
	if err == nil {
		t.Error("Testing wrong options failed: expected error")
	}
}
