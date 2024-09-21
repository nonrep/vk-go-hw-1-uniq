package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Options struct {
	Count      bool
	Duplicate  bool
	Unique     bool
	SkipFields int
	SkipChars  int
	IgnoreCase bool
}

func uniq(input io.Reader, output io.Writer, options Options) error {
	// validation
	if options.Count && options.Duplicate || options.Count && options.Unique || options.Duplicate && options.Unique {
		return errors.New("use only one of this: -c -d -u")
	}
	if options.SkipFields < 0 {
		return errors.New("wrong number of fields to skip")
	}
	if options.SkipChars < 0 {
		return errors.New("wrong number of chars to skip")
	}

	in := bufio.NewScanner(input)
	normal := !(options.Count || options.Duplicate || options.Unique)
	counter := 0
	var prev string
	var prevOriginal string // for output

	for in.Scan() {
		originalString := in.Text() // originalString for output
		txt := originalString
		if options.IgnoreCase {
			txt = strings.ToLower(txt)
		}
		if options.SkipFields > 0 {
			fields := strings.Fields(txt)
			if options.SkipFields >= len(fields) {
				txt = " "
			} else {
				txt = strings.Join(fields[options.SkipFields:], " ")
			}
		}
		if options.SkipChars > 0 {
			if options.SkipChars >= len(txt) {
				txt = " "
			} else {
				txt = txt[options.SkipChars:]
			}
		}

		if txt == prev {
			counter++

			switch {
			case normal:
				continue
			case options.Duplicate && counter == 2:
				fmt.Fprintln(output, originalString)
			}
		} else {
			switch {
			case normal:
				fmt.Fprintln(output, originalString)
			case options.Unique && counter == 1:
				fmt.Fprintln(output, prevOriginal)
			case options.Count && counter != 0:
				fmt.Fprintln(output, strconv.Itoa(counter)+" "+prevOriginal)
			}
			counter = 1
			prev = txt
			prevOriginal = originalString
		}
	}

	if options.Count {
		fmt.Fprintln(output, strconv.Itoa(counter)+" "+prev)
	}

	return nil
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	count := flag.Bool("c", false, "count")
	duplicate := flag.Bool("d", false, "duplicate")
	unique := flag.Bool("u", false, "unique")
	skipFields := flag.Int("f", 0, "skip fields")
	skipChars := flag.Int("s", 0, "skip chars")
	ignoreCase := flag.Bool("i", false, "ignore case")

	flag.Parse()

	options := Options{
		Count:      *count,
		Duplicate:  *duplicate,
		Unique:     *unique,
		SkipFields: *skipFields,
		SkipChars:  *skipChars,
		IgnoreCase: *ignoreCase,
	}

	args := flag.Args()
	in := os.Stdin
	out := os.Stdout
	var err error

	if len(args) > 0 {
		in, err = os.Open(args[0])
		checkError(err)
		defer in.Close()

		if len(args) > 1 {
			out, err = os.Create(args[1])
			checkError(err)
			defer out.Close()
		}
	}

	err = uniq(in, out, options)
	checkError(err)
}
