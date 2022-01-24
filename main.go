package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var (
		isCompact   bool
		isPrettify  bool
		isStringify bool

		out string
		err error
	)

	flag.BoolVar(&isCompact, "c", false, "Compact: Compact JSON input in single line")
	flag.BoolVar(&isPrettify, "p", false, "Prettify: Print prettified JSON input")
	flag.BoolVar(&isStringify, "s", false, "Stringify: Print stringified JSON string")
	flag.Parse()

	if isStringify || isCompact || isPrettify {
		fmt.Println("Invalid argument, -s cannot be used with -c or -p")
		os.Exit(0)
	}

	var in string
	if len(flag.Args()) > 0 {
		in = strings.Join(flag.Args(), "")
	} else {
		scanner := bufio.NewScanner(os.Stdin)

		for scanner.Scan() {
			line := scanner.Text()
			if line == "" {
				break
			} else {
				in += line
			}
		}
	}

	out = in
	if isCompact {
		out, err = Compact(out)
		if err != nil {
			fmt.Printf("Invalid JSON format: %v\n", err)
		}
	} else if isPrettify {
		out, err = Prettify(out)
		if err != nil {
			fmt.Printf("Invalid JSON format: %v\n", err)
		}
	} else if isStringify {
		out, err = Stringify(out)
		if err != nil {
			fmt.Printf("Invalid JSON format: %v\n", err)
		}
	}

	fmt.Println(out)
}

func unquote(s string) string {
	out, err := strconv.Unquote(s)
	if err != nil {
		if err == strconv.ErrSyntax {
			return s
		} else {
			panic(err)
		}
	}
	return out
}

func Compact(s string) (string, error) {
	s = unquote(s)
	buf := &bytes.Buffer{}
	if err := json.Compact(buf, []byte(s)); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func Prettify(s string) (string, error) {
	s = unquote(s)
	buf := &bytes.Buffer{}
	if err := json.Indent(buf, []byte(s), "", "    "); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func Stringify(s string) (string, error) {
	return strings.ReplaceAll(fmt.Sprintf("%q", s), " ", ""), nil
}
