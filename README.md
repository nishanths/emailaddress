# emailaddress

## Godoc

https://pkg.go.dev/github.com/nishanths/emailaddress

Package emailaddress parses and validates emails addresses, using the
standard library's net/mail package. The goal of the package is to never
classify a valid email address as invalid; this is in contrast to the overly
strict email address validation behavior often seen on the Internet.

This package at its core uses the net/mail package for parsing. The net/mail
package appears to classify certain email addresses that are invalid
according to RFCs and Wikipedia as valid. The same behavior is seen from
this package as well.

The net/mail package only provides an interface to parse addresses of the
form "Barry Gibbs <bg@example.com>". This package provides an interaface to
parse and validate standalone email addresses without the name (i.e. of the
form "bg@example.com").

## API

```
func IsValid(email string) bool
func Parse(email string) (localPart, domain string, err error)
```

## Contributing

* If you plan to adpot this package for your project, please consider
  contributing back valid/invalid email address test cases you may already
  have in your project.
