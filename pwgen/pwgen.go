/*
 * Sal's Random Password Generator
 * --------------------------------
 * My very first Go application.
 * It does what it says above.
 * --------------------------------
 * Copyright (c) 2015, Salvatore LaMendola <salvatore@lamendola.me>
 * All rights reserved.
 */

package pwgen

import (
	"crypto/rand"
	"fmt"
	"io"
)

type CharSet int

// Alphanumeric values and symbols+alpha
const (
	AlphaNumeric CharSet = iota
	Symbols
)

func setFor(c CharSet) (string, error) {
	switch c {
	case AlphaNumeric:
		return "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890", nil
	case Symbols:
		a, err := setFor(AlphaNumeric)
		if err != nil {
			return "", err
		}
		return a + "-_!@#$%^&*/\\()_+{}|:<>?=", nil
	default:
		return "", fmt.Errorf("CharSet %d not recognized", c)
	}
}

// Pwgen takes a given length and CharSet and returns a []byte password, or an
// error if the CharSet is not recognized.
func Pwgen(length int, set CharSet) ([]byte, error) {
	password := make([]byte, length)
	entropy := make([]byte, length+(length/4))

	allowedChars, err := setFor(set)
	if err != nil {
		return nil, err
	}

	allowedLength := len(allowedChars)

	// Generate password of the requested length
	io.ReadFull(rand.Reader, entropy)
	for j := 0; j < length; j++ {
		password[j] = allowedChars[entropy[j]%byte(allowedLength)]
	}

	return password, nil
}
