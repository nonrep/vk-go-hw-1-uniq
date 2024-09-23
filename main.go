package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
)

func flagParser() (args []string, options Options) {
	count := flag.Bool("c", false, "count")
	duplicate := flag.Bool("d", false, "duplicate")
	unique := flag.Bool("u", false, "unique")
	skipFields := flag.Int("f", 0, "skip fields")
	skipChars := flag.Int("s", 0, "skip chars")
	ignoreCase := flag.Bool("i", false, "ignore case")

	flag.Parse()

	options = Options{
		Count:      *count,
		Duplicate:  *duplicate,
		Unique:     *unique,
		SkipFields: *skipFields,
		SkipChars:  *skipChars,
		IgnoreCase: *ignoreCase,
	}

	return flag.Args(), options
}

func argsParser(args []string) (in, out *os.File, err error) {
	in = os.Stdin
	out = os.Stdout

	// если 2 аргумента, то необходимо выполнить и case 2, и case 1, поэтому fallthrough
	switch len(args) {
	case 0:
		// ничего не делаем
	case 2:
		out, err = os.Create(args[1])
		if err != nil {
			fmt.Println(err)
			return in, out, err
		}
		fallthrough
	case 1:
		in, err = os.Open(args[0])
		if err != nil {
			return in, out, err
		}
	default:
		return in, out, errors.New("too many arguments")
	}
	return in, out, nil
}

func getStrings(input io.Reader) (strings []string) {
	in := bufio.NewScanner(input)
	for in.Scan() {
		strings = append(strings, in.Text())
	}
	return strings
}

func main() {
	args, options := flagParser()

	in, out, err := argsParser(args)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer in.Close()
	defer out.Close()

	strings := getStrings(in)

	result, err := uniq(strings, options)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, str := range result {
		fmt.Fprintln(out, str)
	}

}
