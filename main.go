package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/nonrep/go-homework-1-uniq/uniq"
)

// flagParser извлекает определенные флаги командной строки, возвращает набор флагов и неопознанные параметры командой строки.
func flagParser() (args []string, options uniq.Options) {
	count := flag.Bool("c", false, "count")
	duplicate := flag.Bool("d", false, "duplicate")
	unique := flag.Bool("u", false, "unique")
	skipFields := flag.Int("f", 0, "skip fields")
	skipChars := flag.Int("s", 0, "skip chars")
	ignoreCase := flag.Bool("i", false, "ignore case")

	flag.Parse()

	options = uniq.Options{
		Count:      *count,
		Duplicate:  *duplicate,
		Unique:     *unique,
		SkipFields: *skipFields,
		SkipChars:  *skipChars,
		IgnoreCase: *ignoreCase,
	}

	return flag.Args(), options
}

// argsParser считывает неопознанные аргументы командой строки после выполнения flagParser, ожидает текстовые файлы.
func argsParser(args []string) (in, out *os.File, err error) {
	in = os.Stdin
	out = os.Stdout

	// Если 2 аргумента, то необходимо выполнить и case 2, и case 1, поэтому fallthrough.
	switch len(args) {
	case 0:
		// Ничего не делаем.
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

// getStrings считывает строки из входного потока, возвращает слайс строк.
func getStrings(input io.Reader) (strings []string, err error) {
	in := bufio.NewScanner(input)
	for in.Scan() {
		strings = append(strings, in.Text())
	}
	if err := in.Err(); err != nil {
		return strings, err
	}
	return strings, nil
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

	strings, err := getStrings(in)
	if err != nil {
		fmt.Println(err)
		return
	}

	result, err := uniq.Uniq(strings, options)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, str := range result {
		fmt.Fprintln(out, str)
	}

}
