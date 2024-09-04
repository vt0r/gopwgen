/*
 * Sal's Random Password Generator (GoPwgen)
 * --------------------------------
 * My very first Go application.
 * It does what it says above.
 * --------------------------------
 * Copyright (c) 2015-2024, Salvatore LaMendola <salvatore@lamendo.la>
 * All rights reserved.
 */

package main

import (
	"bytes"
	"crypto/rand"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"

	"github.com/sethvargo/go-diceware/diceware"
)

// Assign all the acceptable arguments and their default values
var (
	flagSymbols     = flag.Bool("s", false, "Alphanumeric + symbols (NOT FOR MYSQL!)")
	flagAlpha       = flag.Bool("a", false, "Alphanumeric only")
	flagDiceware    = flag.Bool("d", false, "Diceware passphrase (Choose random words from a list)")
	flagHexadecimal = flag.Bool("H", false, "Hexadecimal only (abcdef0123456789)")
	flagPhpMyAdmin  = flag.Bool("p", false, "Generate phpMyAdmin Blowfish secret (for cookie auth)")
	flagWordPress   = flag.Bool("w", false, "Generate WordPress encryption salts for use in wp-config.php")
	flagVersion     = flag.Bool("v", false, "Print the version number and exit")

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
const hexadecimal = "abcdef0123456789"
const defaultlen = 19
const defaultnum = 1
const defaultDwLen = 6
const version = "2.2.1"

func myUsage() {
	fmt.Printf("Usage: %s [OPTION] [length | num dice words] [num pwds/phrases]\n\nOptions:\n", os.Args[0])
	flag.PrintDefaults()
}

func myVersion() {
	fmt.Printf("GoPwgen version %s\n", version)
}

func diceWareGen(dwlen int) {
	list, err := diceware.Generate(dwlen)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", strings.Join(list, "-"))
}

func main() {
	flag.Usage = myUsage
	flag.Parse()

	// Either set len/num using provided values or fallback to defaults
	allowed := symbols
	pwlen, err1 := strconv.Atoi(flag.Arg(0))
	dwlen := pwlen
	if err1 != nil {
		pwlen = defaultlen
		dwlen = defaultDwLen
	}
	numPws, err2 := strconv.Atoi(flag.Arg(1))
	if err2 != nil {
		numPws = defaultnum
	}

	pwStringer := func(n int, s string) string { return string(pwgen(n, s)) }

	// Option validation
	switch {
	case *flagVersion:
		myVersion()
		os.Exit(0)
	case *flagSymbols: // This is the default assigned value
	case *flagAlpha:
		allowed = alphanumeric
	case *flagDiceware:
		for i := 0; i < numPws; i++ {
			diceWareGen(dwlen)
		}
		os.Exit(0)
	case *flagHexadecimal:
		allowed = hexadecimal
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
	for i := 0; i < numPws; i++ {
		outputs[i] = pwStringer(pwlen, allowed)
	}
	fmt.Println(strings.Join(outputs, "\n"))
}

func pwgen(length int, allowedChars string) []byte {
	// Create the password string and associated randomness
	password := make([]byte, length)
	entropy := make([]byte, length+(length/4))
	allowedLength := len(allowedChars)

	// Generate password of the requested length
	_, err := io.ReadFull(rand.Reader, entropy)
	if err != nil {
		fmt.Println("Error: ", err)
	}
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
