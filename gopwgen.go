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
	"crypto/rand"
	"flag"
	"fmt"
	"io"
	"os"
)

// Assign all the acceptable arguments and their default values
var (
	flagSymbols    = flag.Bool("s", false, "Alphanumeric + symbols (NOT FOR MYSQL!)")
	flagAlpha      = flag.Bool("a", false, "Alphanumeric only")
	flagPhpMyAdmin = flag.Bool("p", false, "Generate phpMyAdmin Blowfish secret (for cookie auth)")
	flagWordPress  = flag.Bool("w", false, "Generate WordPress encryption salts for use in wp-config.php")
	flagLength     = flag.Int("l", 19, "Length of generated password(s)")
	flagNumber     = flag.Int("n", 1, "Number of generated password(s)")
)

// Alphanumeric values and symbols+alpha
const alphanumeric = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
const symbols = alphanumeric + "-_!@#$%^&*/\\()_+{}|:<>?="

func main() {
	flag.Parse()

	// No option specified
	if !*flagAlpha && !*flagSymbols && !*flagPhpMyAdmin && !*flagWordPress {
		fmt.Printf("ERROR: No option selected.\n\n")
		flag.Usage()
		os.Exit(1)
	}

	// Generate alphanumeric password(s)
	if *flagAlpha {
		for i := 0; i < *flagNumber; i++ {
			password := pwgen(*flagLength, alphanumeric)
			fmt.Println(string(password))
		}
	}

	// Generate alpha/symbols password(s)
	if *flagSymbols {
		for i := 0; i < *flagNumber; i++ {
			password := pwgen(*flagLength, symbols)
			fmt.Println(string(password))
		}
	}

	// Generate phpMyAdmin blowfish secret
	if *flagPhpMyAdmin {
		password := pwgen(64, symbols)
		fmt.Println(string(password))
	}

	// Generate WordPress encryption secrets
	if *flagWordPress {
		fmt.Printf("define('AUTH_KEY',\t\t'%s');\n", pwgen(64, symbols))
		fmt.Printf("define('SECURE_AUTH_KEY',\t'%s');\n", pwgen(64, symbols))
		fmt.Printf("define('LOGGED_IN_KEY',\t\t'%s');\n", pwgen(64, symbols))
		fmt.Printf("define('NONCE_KEY',\t\t'%s');\n", pwgen(64, symbols))
		fmt.Printf("define('AUTH_SALT',\t\t'%s');\n", pwgen(64, symbols))
		fmt.Printf("define('SECURE_AUTH_SALT',\t'%s');\n", pwgen(64, symbols))
		fmt.Printf("define('LOGGED_IN_SALT',\t'%s');\n", pwgen(64, symbols))
		fmt.Printf("define('NONCE_SALT',\t\t'%s');\n", pwgen(64, symbols))
	}

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
