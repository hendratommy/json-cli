package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	var (
		isStringify bool
		isPrettify  bool

		out string
		err error
	)

	flag.BoolVar(&isStringify, "s", false, "Print JSON in single line")
	flag.BoolVar(&isStringify, "stringify", false, "Print JSON in single line")
	flag.BoolVar(&isPrettify, "p", false, "Print the prettified JSON input")
	flag.BoolVar(&isPrettify, "prettify", false, "Print the prettified JSON input")
	flag.Parse()

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
	if isStringify {
		out, err = Stringify(out)
		if err != nil {
			fmt.Printf("Invalid JSON format: %v\n", err)
		}
	}
	if isPrettify {
		out, err = Prettify(out)
		if err != nil {
			fmt.Printf("Invalid JSON format: %v\n", err)
		}
	}

	fmt.Println(out)
}

func Stringify(s string) (string, error) {
	buf := &bytes.Buffer{}
	if err := json.Compact(buf, []byte(s)); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func Prettify(s string) (string, error) {
	buf := &bytes.Buffer{}
	if err := json.Indent(buf, []byte(s), "", "    "); err != nil {
		return "", err
	}
	return buf.String(), nil
}
