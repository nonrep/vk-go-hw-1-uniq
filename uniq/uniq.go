package uniq

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

// validate проверяет наличие недопустимых наборов и значений флагов.
func validate(options Options) error {
	if options.Count && options.Duplicate || options.Count && options.Unique || options.Duplicate && options.Unique {
		return errors.New("use only one of this: -c -d -u")
	}
	if options.SkipFields < 0 {
		return errors.New("wrong number of fields to skip")
	}
	if options.SkipChars < 0 {
		return errors.New("wrong number of chars to skip")
	}
	return nil
}

// applyOptions изменяет строки для корректного сравнения внутри Uniq, если были переданы следующие флаги: -s -i -c.
func applyOptions(string string, options Options) string {
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

// Uniq преобразует полученный слайс строк в слайс строк в соответствии с параметрами options.
func Uniq(stringsSlice []string, options Options) (result []string, err error) {
	if err := validate(options); err != nil {
		return result, err
	}

	normal := !(options.Count || options.Duplicate || options.Unique)

	var (
		i            int // Итератор stringCounts.
		prevString   string
		stringCounts []Pair
	)

	for _, str := range stringsSlice {
		stringToCompare := applyOptions(str, options)
		if stringToCompare == prevString {
			stringCounts[i].count++
		} else {
			stringCounts = append(stringCounts, Pair{string: str, count: 1})
			prevString = stringToCompare
			i = len(stringCounts) - 1
		}
	}

	for _, pair := range stringCounts {
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
