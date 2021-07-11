# emailaddress

https://pkg.go.dev/github.com/nishanths/emailaddress

The goal of the package is to never classify a valid email address as invalid;
this is in contrast to the overly strict email address validation behavior often
seen on the Internet.

It uses net/mail underneath.

## API

```
func IsValid(email string) bool
func Parse(email string) (localPart, domain string, err error)
```

## Contributing

* If you plan to adopt this package for your project, please consider
  contributing back valid/invalid email address test cases you may already
  have in your project.
