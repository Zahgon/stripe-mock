package parser

import (
	"github.com/stripe/stripe-mock/param/form"
)

//
// Public functions
//

// ParseFormString parses a form-encoded body or query into a set of key/value
// pairs. It differs from url.ParseQuery in that because it produces a slice
// instead of a map, order can be preserved. This is key to properly decoding
// "Rack-style" form encoding.
//
// Implementation modified from: https://github.com/deoxxa/urlqp
func ParseFormString(s string) (form.Values, error) {
	_ = "STUB: not implemented"
	return *new(form.Values), nil
}

// Split this raw form value into two parts, at the first `=`

// Set a default for the value. Empty seems reasonable.

// If `b` has more than one element, that means the second one will be the
// parameter value, so grab it.
