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
	"crypto/rand"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"text/tabwriter"
)

// Assign all the acceptable arguments and their default values
var (
	flagSymbols    = flag.Bool("s", false, "Alphanumeric + symbols (NOT FOR MYSQL!)")
	flagAlpha      = flag.Bool("a", false, "Alphanumeric only")
	flagPhpMyAdmin = flag.Bool("p", false, "Generate phpMyAdmin Blowfish secret (for cookie auth)")
	flagWordPress  = flag.Bool("w", false, "Generate WordPress encryption salts for use in wp-config.php")
	flagLength     = flag.Int("l", 19, "Length of generated password(s)")
	flagNumber     = flag.Int("n", 1, "Number of generated password(s)")

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
)

// Alphanumeric values and symbols+alpha
const alphanumeric = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
const symbols = alphanumeric + "-_!@#$%^&*/\\()_+{}|:<>?="

func main() {
	flag.Parse()

	// No option specified
	allowed, numPws, pwlen := symbols, *flagNumber, *flagLength

	pwStringer := func(n int, s string) string { return string(pwgen(n, s)) }

	switch {
	case *flagSymbols: // Already set up this way
	case *flagAlpha:
		allowed = alphanumeric
	case *flagPhpMyAdmin:
		pwlen = 64
	case *flagWordPress:
		pwlen = 64
		w := new(tabwriter.Writer)
		var b bytes.Buffer
		w.Init(&b, 26, 1, 0, ' ', 0)
		pwStringer = phpKeysPwgen(w, &b)
	default:
		fmt.Printf("ERROR: No option selected.\n\n")
		flag.Usage()
		os.Exit(1)
	}

	outputs := make([]string, numPws)
	for i := 0; i < numPws; i++ {
		outputs[i] = pwStringer(pwlen, allowed)
	}
	fmt.Print(strings.Join(outputs, "\n"))
}

func pwgen(length int, allowedChars string) []byte {
	// Create the password string and associated randomness
	password := make([]byte, length)
	entropy := make([]byte, length+(length/4))
	allowedLength := len(allowedChars)

	// Generate password of the requested length
	io.ReadFull(rand.Reader, entropy)
	for j := 0; j < length; j++ {
		password[j] = allowedChars[entropy[j]%byte(allowedLength)]
	}
	return password
}

type flushableWriter interface {
	io.Writer
	Flush() error
}

func phpKeysPwgen(w flushableWriter, b *bytes.Buffer) func(int, string) string {
	return func(n int, s string) string {
		lines := make([]string, len(phpKeys)+1)
		for i, key := range phpKeys {
			fmt.Fprintf(w, "define('%s',\t'%s');", key, string(pwgen(n, s)))
			w.Flush()

			lines[i] = b.String()
			b.Reset()
		}
		return strings.Join(lines, "\n")
	}
}
