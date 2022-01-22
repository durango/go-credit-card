// Package creditcard provides methods for validating credit cards
package creditcard

import (
	"errors"
	"time"
)

var timeNowCaller = time.Now

var (
	errCardHasExpired          = errors.New("credit card has expired")
	errCardNumberNotLongEnough = errors.New("credit card number is not long enough")
	errInvalidCVV              = errors.New("invalid CVV")
	errInvalidCardNumber       = errors.New("invalid credit card number")
	errInvalidMonth            = errors.New("invalid month")
	errInvalidYear             = errors.New("invalid year")
	errTestNumbersNotAllowed   = errors.New("test numbers are not allowed")
	errUnkownCardMethod        = errors.New("unknown credit card method")
)

// Luhner interface represents a credit card validator tool.
type Luhner interface {
	LastFour() (string, error)
	LastFourDigits() (string, error)
	Wipe()
	Validate(allowTestNumbers ...bool) error
	ValidateExpiration() error
	ValidateCVV() error
	Method() error
	MethodValidate() (Company, error)
	ValidateNumber() bool
}

func resetMocks() {
	timeNowCaller = time.Now
}

func matchesValue(number int, numbers []int) bool {
	for _, v := range numbers {
		if v == number {
			return true
		}
	}
	return false
}

func isInBetween(n, min, max int) bool {
	return n >= min && n <= max
}
