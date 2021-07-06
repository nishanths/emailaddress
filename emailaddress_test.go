package emailaddress

import "testing"

func TestCheck(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		testcases := []struct {
			email, localpart, domain string
		}{
			// from Wikipedia: https://en.wikipedia.org/wiki/Email_address#Examples
			{"simple@example.com", "simple", "example.com"},
			{"very.common@example.com", "very.common", "example.com"},
			{"disposable.style.email.with+symbol@example.com", "disposable.style.email.with+symbol", "example.com"},
			{"other.email-with-hyphen@example.com", "other.email-with-hyphen", "example.com"},
			{"fully-qualified-domain@example.com", "fully-qualified-domain", "example.com"},
			{"user.name+tag+sorting@example.com", "user.name+tag+sorting", "example.com"},
			{"x@example.com", "x", "example.com"},
			{"example-indeed@strange-example.com", "example-indeed", "strange-example.com"},
			{"test/test@test.com", "test/test", "test.com"},
			{"admin@mailserver1", "admin", "mailserver1"},
			{"example@s.example", "example", "s.example"},
			{`" "@example.org`, `" "`, "example.org"},
			{`"john..doe"@example.org`, `"john..doe"`, "example.org"},
			{"mailhost!username@example.org", "mailhost!username", "example.org"},
			{"user%example.com@example.org", "user%example.com", "example.org"},
			{"user-@example.org", "user-", "example.org"},

			// Listed as invalid at https://en.wikipedia.org/wiki/Email_address#Examples.
			// But considered valid by this package (and by net/mail).
			{"1234567890123456789012345678901234567890123456789012345678901234+x@example.com", "1234567890123456789012345678901234567890123456789012345678901234+x", "example.com"}, // local-part is longer than 64 characters
			{"i_like_underscore@but_its_not_allowed_in_this_part.example.com", "i_like_underscore", "but_its_not_allowed_in_this_part.example.com"},                                 // Underscore is not allowed in domain part

			// Invalid based on descriptions at https://en.wikipedia.org/wiki/Email_address.
			// But considered valid by this package (and by net/mail).
			{"simple@-example.com", "simple", "-example.com"},
			{"simple@example.com-", "simple", "example.com-"},
			{"simple@-example.com-", "simple", "-example.com-"},

			{"under_score@example.com", "under_score", "example.com"},
			{"simple@1234567890123456789012345678901234567890123456789012345678901234xx.com", "simple", "1234567890123456789012345678901234567890123456789012345678901234xx.com"},
		}

		for _, tt := range testcases {
			t.Run(tt.email, func(t *testing.T) {
				gotlocal, gotdomain, err := Parse(tt.email)
				requireNoError(t, err)
				requireEqual(t, tt.localpart, gotlocal)
				requireEqual(t, tt.domain, gotdomain)
			})
		}
	})

	const (
		errWhitespace    = "white space around email address"
		errAngleBrackets = "angle brackets around email address"
		errStandalone    = "not standalone email address"
	)

	t.Run("invalid", func(t *testing.T) {
		testcases := []struct {
			email string
			err   string
		}{
			// from Wikipedia: https://en.wikipedia.org/wiki/Email_address#Examples
			{"Abc.example.com", "failed to parse address: mail: missing '@' or angle-addr"},                         // (no @ character
			{"A@b@c@example.com", "failed to parse address: mail: expected single address, got \"@c@example.com\""}, // only one @ is allowed outside quotation marks
			{`a"b(c)d,e:f;g<h>i[j\k]l@example.com`, "failed to parse address: mail: missing '@' or angle-addr"},     // none of the special characters in this local-part are allowed outside quotation marks
			{`just"not"right@example.com`, "failed to parse address: mail: no angle-addr"},                          // quoted strings must be dot separated or the only element making up the local-part
			{`this is"not\allowed@example.com`, "failed to parse address: mail: no angle-addr"},                     // spaces, quotes, and backslashes may only exist when within quoted strings and preceded by a backslash
			{`this\ still\"not\\allowed@example.com`, "failed to parse address: mail: no angle-addr"},               // even if escaped (preceded by a backslash), spaces, quotes, and backslashes must still be contained by quotes
			{"QA[icon]CHOCOLATE[icon]@test.com", "failed to parse address: mail: no angle-addr"},                    // icon characters

			// white space
			{"", "failed to parse address: mail: no address"},
			{" ", errWhitespace},
			{" simple@example.com", errWhitespace},
			{"simple@example.com ", errWhitespace},
			{" simple@example.com ", errWhitespace},

			// angle brackets
			{"<simple@example.com>", errAngleBrackets},
			{"<<simple@example.com>>", errAngleBrackets},
			{"<simple@example.com", "failed to parse address: mail: unclosed angle-addr"},
			{"simple@example.com>", "failed to parse address: mail: expected single address, got \">\""},

			// name + address
			{"Barry Gibbs <bg@example.com>", errStandalone},
			{"Barry Gibbs", "failed to parse address: mail: no angle-addr"},
			{"Barry Gibbs <>", "failed to parse address: mail: invalid string"},

			// no domain
			{"gb@", "failed to parse address: mail: no angle-addr"},
		}

		for _, tt := range testcases {
			t.Run(tt.email, func(t *testing.T) {
				_, _, err := Parse(tt.email)
				requireError(t, err)
				requireEqual(t, tt.err, err.Error())
			})
		}
	})
}

func requireNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("expected nil error, got: %q", err)
	}
}

func requireError(t *testing.T, err error) {
	t.Helper()
	if err == nil {
		t.Fatal("expected non-nil error, got nil")
	}
}

func requireEqual(t *testing.T, expect, got string) {
	t.Helper()
	if expect != got {
		t.Fatalf("expected: %q, got: %q", expect, got)
	}
}
