// Package emailaddress parses and validates emails addresses, using the standard
// library's net/mail package. The goal of the package is to never
// classify a valid email address as invalid; this is in contrast to the overly strict
// email address validation behavior often seen on the Internet.
//
// This package at its core uses the net/mail package for parsing. The
// net/mail package appears to classify certain email addresses that are invalid
// according to RFCs and Wikipedia as valid. The same behavior is seen from this
// package as well.
//
// The net/mail package only provides an interface to parse addresses of the
// form "Barry Gibbs <bg@example.com>". This package provides an interaface to
// parse and validate standalone email addresses without the name (i.e. of
// the form "bg@example.com").
package emailaddress

import (
	"errors"
	"fmt"
	"net/mail"
	"strings"
)

// Parse parses an email address into its local-part and domain.
// If the error is nil, the format of the email address can be considered
// valid.
func Parse(email string) (localPart, domain string, err error) {
	if len(strings.TrimSpace(email)) != len(email) {
		return "", "", errors.New("white space around email address")
	}

	if strings.HasPrefix(email, "<") && strings.HasSuffix(email, ">") {
		return "", "", errors.New("angle brackets around email address")
	}

	addr, err := mail.ParseAddress(email)
	if err != nil {
		return "", "", fmt.Errorf("failed to parse address: %w", err)
	}

	if addr.Name != "" {
		return "", "", errors.New("not standalone email address")
	}

	idx := strings.LastIndexByte(email, '@')
	if idx == -1 {
		panic("emailaddress: unexpectedly failed to find '@' in email address " + email)
	}

	localPart = email[:idx]
	domain = email[idx+1:]
	return localPart, domain, nil
}

// IsValid returns whether the email address is of valid format.
// It is a shorthand for the piece of code:
//     _, _, err := Parse(email)
//     return err != nil
func IsValid(email string) bool {
	_, _, err := Parse(email)
	return err != nil
}