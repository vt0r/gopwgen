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

/*
 * Proper flag variable assignment and usage info
 * Thanks to Jacques Fuentes @jpfuentes2 for the code review!
 */
var (
	flagSymbols    = flag.Bool("s", false, "Alphanumeric + symbols (NOT FOR MYSQL!)")
	flagAlpha      = flag.Bool("a", false, "Alphanumeric only")
	flagPhpMyAdmin = flag.Bool("p", false, "Generate phpMyAdmin Blowfish secret (for cookie auth)")
	flagWordPress  = flag.Bool("w", false, "Generate WordPress encryption salts for use in wp-config.php")
	flagLength     = flag.Int("l", 19, "Length of generated password(s)")
	flagNumber     = flag.Int("n", 1, "Number of generated password(s)")
)

// Alphanumeric values and symbols+alpha
const alphanumeric = "abcdefghjkmnpqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ23456789"
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
			password := pwgen(*flagLength, false)
			fmt.Println(string(password))
		}
	}

	// Generate alpha/symbols password(s)
	if *flagSymbols {
		for i := 0; i < *flagNumber; i++ {
			password := pwgen(*flagLength, true)
			fmt.Println(string(password))
		}
	}

	// Generate phpMyAdmin blowfish secret
	if *flagPhpMyAdmin {
		password := pwgen(64, true)
		fmt.Println(string(password))
	}

	// Generate WordPress encryption secrets
	if *flagWordPress {
		password1 := pwgen(64, true)
		password2 := pwgen(64, true)
		password3 := pwgen(64, true)
		password4 := pwgen(64, true)
		password5 := pwgen(64, true)
		password6 := pwgen(64, true)
		password7 := pwgen(64, true)
		password8 := pwgen(64, true)

		fmt.Printf("define('AUTH_KEY',\t\t'%s');\n", password1)
		fmt.Printf("define('SECURE_AUTH_KEY',\t'%s');\n", password2)
		fmt.Printf("define('LOGGED_IN_KEY',\t\t'%s');\n", password3)
		fmt.Printf("define('NONCE_KEY',\t\t'%s');\n", password4)
		fmt.Printf("define('AUTH_SALT',\t\t'%s');\n", password5)
		fmt.Printf("define('SECURE_AUTH_SALT',\t'%s');\n", password6)
		fmt.Printf("define('LOGGED_IN_SALT',\t'%s');\n", password7)
		fmt.Printf("define('NONCE_SALT',\t\t'%s');\n", password8)

	}

}

func pwgen(length int, symbolsenabled bool) []byte {
	// Create the password string and associated randomness
	password := make([]byte, length)
	entropy := make([]byte, length+(length/4))

	// Generate password(s) of the requested length/count
	io.ReadFull(rand.Reader, entropy)
	for j := 0; j < length; j++ {
		if symbolsenabled {
			password[j] = symbols[entropy[j]%byte(len(symbols))]
		} else {
			password[j] = alphanumeric[entropy[j]%byte(len(alphanumeric))]
		}
	}
	return password
}
