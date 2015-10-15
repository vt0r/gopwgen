/*
 * Sal's Random Password Generator
 * --------------------------------
 * My very first Go application.
 * It does what it says above.
 * --------------------------------
 * Copyright (c) 2015, Salvatore LaMendola <salvatore@lamendola.me>
 * All rights reserved.
 */

package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"

	"github.com/vt0r/gopwgen/pwgen"
)

// Assign all the acceptable arguments and their default values
var (
	flagSymbols    = flag.Bool("s", false, "Alphanumeric + symbols (NOT FOR MYSQL!)")
	flagAlpha      = flag.Bool("a", false, "Alphanumeric only")
	flagPhpMyAdmin = flag.Bool("p", false, "Generate phpMyAdmin Blowfish secret (for cookie auth)")
	flagWordPress  = flag.Bool("w", false, "Generate WordPress encryption salts for use in wp-config.php")

	phpKeys = [...]string{
		"AUTH_KEY",
		"SECURE_AUTH_KEY",
		"LOGGED_IN_KEY",
		"NONCE_KEY",
		"AUTH_SALT",
		"SECURE_AUTH_SALT",
		"LOGGED_IN_SALT",
		"NONCE_SALT",
	}

	defaultLen = 19
	defaultNum = 1
)

func myUsage() {
	fmt.Printf("Usage: %s [OPTION] [length] [number]\n\nOptions:\n", os.Args[0])
	flag.PrintDefaults()
}

func main() {
	flag.Usage = myUsage
	flag.Parse()

	// Either set len/num using provided values or fallback to defaults
	allowed := pwgen.Symbols
	pwlen, err1 := strconv.Atoi(flag.Arg(0))
	if err1 != nil {
		pwlen = defaultLen
	}
	numPws, err2 := strconv.Atoi(flag.Arg(1))
	if err2 != nil {
		numPws = defaultNum
	}

	pwStringer := func(n int, c pwgen.CharSet) (string, error) {
		s, err := pwgen.Pwgen(n, c)
		if err != nil {
			return "", err
		}
		return string(s), nil
	}

	// Option validation
	switch {
	case *flagSymbols: // This is the default assigned value
	case *flagAlpha:
		allowed = pwgen.AlphaNumeric
	case *flagPhpMyAdmin:
		pwlen = 64
	case *flagWordPress:
		pwlen = 64
		w := new(tabwriter.Writer)
		var b bytes.Buffer
		w.Init(&b, 26, 1, 0, ' ', 0)
		pwStringer = phpKeysPwgen(w, &b)
	}

	outputs := make([]string, numPws)

	var (
		pw  string
		err error
	)

	for i := 0; i < numPws; i++ {
		pw, err = pwStringer(pwlen, allowed)
		if err != nil {
			log.Fatalf("error while generating password %d: %v", i, err)
			outputs[i] = pw
		}
	}

	fmt.Print(strings.Join(outputs, "\n"))
}

type flushableWriter interface {
	io.Writer
	Flush() error
}

func phpKeysPwgen(w flushableWriter, b *bytes.Buffer) func(int, pwgen.CharSet) (string, error) {
	return func(n int, c pwgen.CharSet) (string, error) {
		var (
			pw  []byte
			err error

			lines = make([]string, len(phpKeys)+1)
		)

		for i, key := range phpKeys {
			pw, err = pwgen.Pwgen(n, c)
			if err != nil {
				return "", err
			}

			fmt.Fprintf(w, "define('%s',\t'%s');", key, pw)
			w.Flush()

			lines[i] = b.String()
			b.Reset()
		}
		return strings.Join(lines, "\n"), nil
	}
}
