package main

import (
	"errors"
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

type Pair struct {
	string string
	count  int
}

// для сравнения если есть -s -i -c
func changeString(string string, options Options) string {
	if options.IgnoreCase {
		string = strings.ToLower(string)
	}
	if options.SkipFields > 0 {
		fields := strings.Fields(string)
		if options.SkipFields >= len(fields) {
			string = " "
		} else {
			string = strings.Join(fields[options.SkipFields:], " ")
		}
	}
	if options.SkipChars > 0 {
		if options.SkipChars >= len(string) {
			string = " "
		} else {
			string = string[options.SkipChars:]
		}
	}
	return string
}

func uniq(stringsSlice []string, options Options) (result []string, err error) {
	// validation
	if options.Count && options.Duplicate || options.Count && options.Unique || options.Duplicate && options.Unique {
		return result, errors.New("use only one of this: -c -d -u")
	}
	if options.SkipFields < 0 {
		return result, errors.New("wrong number of fields to skip")
	}
	if options.SkipChars < 0 {
		return result, errors.New("wrong number of chars to skip")
	}

	normal := !(options.Count || options.Duplicate || options.Unique)

	var (
		i          int
		prevString string
		counter    []Pair
	)

	for _, str := range stringsSlice {
		stringToCompare := changeString(str, options)
		if stringToCompare == prevString {
			counter[i].count++
		} else {
			counter = append(counter, Pair{string: str, count: 1})
			prevString = stringToCompare
			i = len(counter) - 1
		}
	}

	for _, pair := range counter {
		switch {
		case normal:
			result = append(result, pair.string)
		case options.Count:
			result = append(result, strconv.Itoa(pair.count)+" "+pair.string)
		case options.Duplicate:
			if pair.count > 1 {
				result = append(result, pair.string)
			}
		case options.Unique:
			if pair.count == 1 {
				result = append(result, pair.string)
			}
		}
	}

	return result, nil
}
