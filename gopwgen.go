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
	"strconv"
	"strings"
	"text/tabwriter"
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
)

// Alphanumeric values and symbols+alpha / default length and number of passwords
const alphanumeric = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
const symbols = alphanumeric + "-_!@#$%^&*/\\()_+{}|:<>?="
const defaultlen = 19
const defaultnum = 1

func main() {
	flag.Parse()

	// Override the default (ugly) usage output
	flag.Usage = func() {
		fmt.Printf("Sal's Random Password Generator\n")
		fmt.Printf("-------------------------------\n")
		fmt.Printf("Usage: %s <OPTION> [length] [number] (length and number optional)\n\n", os.Args[0])
		fmt.Printf("OPTIONS (MUST SPECIFY ONE!):\n")
		fmt.Printf("-s               Add symbols to output (NOT FOR MYSQL!)\n")
		fmt.Printf("-a               Alphanumeric only\n")
		fmt.Printf("-p               Generate phpMyAdmin Blowfish secret (for cookie auth)\n")
		fmt.Printf("-w               Generate Wordpress encryption keys (wp-config.php)\n")
		fmt.Printf("-h               Display this usage information\n\n")
		fmt.Printf("If no length or number are defined, a default length of %d and number of %d will be used.\n\n", defaultlen, defaultnum)
	}

	// Either set len/num using provided values or fallback to defaults
	allowed := symbols
	pwlen, err1 := strconv.Atoi(flag.Arg(0))
	if err1 != nil {
		pwlen = defaultlen
	}
	numPws, err2 := strconv.Atoi(flag.Arg(1))
	if err2 != nil {
		numPws = defaultnum
	}

	pwStringer := func(n int, s string) string { return string(pwgen(n, s)) }

	// Option validation
	switch {
	case *flagSymbols: // This is the default assigned value
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
		fmt.Printf("ERROR: At least one option must be selected.\n\n")
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
