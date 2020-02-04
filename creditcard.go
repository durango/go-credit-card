// Package creditcard provides methods for validating credit cards
package creditcard

import (
	"errors"
	"strconv"
	"time"
)

// Card holds generic information about the credit card
type Card struct {
	Number, Cvv, Month, Year string
	Company                  Company
}

// Company holds a short and long names of who has issued the credit card
type Company struct {
	Short, Long string
}

type digits [6]int

// At returns the digits from the start to the given length
func (d *digits) At(i int) int {
	return d[i-1]
}

// LastFour returns the last four digits of the credit card's number
func (c *Card) LastFour() (string, error) {
	if len(c.Number) < 4 {
		return "", errors.New("Credit card number is not long enough")
	}

	return c.Number[len(c.Number)-4 : len(c.Number)], nil
}

// LastFourDigits as an alias for LastFour
func (c *Card) LastFourDigits() (string, error) {
	return c.LastFour()
}

// Wipe returns the credit card with false/nullified/generic information
func (c *Card) Wipe() {
	c.Cvv, c.Number, c.Month, c.Year = "0000", "0000000000000000", "01", "1970"
}

// Validate returns nil or an error describing why the credit card didn't validate
// this method checks for expiration date, CCV/CVV and the credit card's numbers.
// For allowing test cards to go through, simply pass true (bool) as the first argument
func (c *Card) Validate(allowTestNumbers ...bool) error {
	err := c.ValidateExpiration()
	if err != nil {
		return err
	}

	err = c.ValidateCVV()
	if err != nil {
		return err
	}

	switch c.Number {
	// test cards: https://stripe.com/docs/testing
	case "4242424242424242",
		"4012888888881881",
		"4000056655665556",
		"5555555555554444",
		"5200828282828210",
		"5105105105105100",
		"378282246310005",
		"371449635398431",
		"6011111111111117",
		"6011000990139424",
		"30569309025904",
		"38520000023237",
		"3530111333300000",
		"3566002020360505",
		"4111111111111111",
		"4916909992637469",
		"4000111111111115",
		"2223000048400011",
		"6035227716427021":
		if len(allowTestNumbers) > 0 && allowTestNumbers[0] {
			return nil
		}

		return errors.New("Test numbers are not allowed")
	}

	valid := c.ValidateNumber()

	if !valid {
		return errors.New("Invalid credit card number")
	}

	return nil
}

// validates the credit card's expiration date
func (c *Card) ValidateExpiration() error {
	var year, month int
	var err error

	if len(c.Year) < 3 {
		year, err = strconv.Atoi(strconv.Itoa(time.Now().UTC().Year())[:2] + c.Year)
		if err != nil {
			return errors.New("Invalid year")
		}
	} else {
		year, err = strconv.Atoi(c.Year)
		if err != nil {
			return errors.New("Invalid year")
		}
	}

	month, err = strconv.Atoi(c.Month)
	if err != nil {
		return errors.New("Invalid month")
	}

	if month < 1 || 12 < month {
		return errors.New("Invalid month")
	}

	if year < time.Now().UTC().Year() {
		return errors.New("Credit card has expired")
	}

	if year == time.Now().UTC().Year() && month < int(time.Now().UTC().Month()) {
		return errors.New("Credit card has expired")
	}

	return nil
}

// validates the length of the card's CVV value
func (c *Card) ValidateCVV() error {
	if len(c.Cvv) < 3 || len(c.Cvv) > 4 {
		return errors.New("Invalid CVV")
	}

	return nil
}

// Method returns an error from MethodValidate() or returns the
// credit card with it's company / issuer attached to it
func (c *Card) Method() error {
	company, err := c.MethodValidate()

	if err != nil {
		return err
	}

	c.Company = company
	return nil
}

// MethodValidate adds/checks/verifies the credit card's company / issuer
func (c *Card) MethodValidate() (Company, error) {
	var err error
	ccLen := len(c.Number)
	ccDigits := digits{}

	for i := 0; i < 6; i++ {
		if i < ccLen {
			ccDigits[i], err = strconv.Atoi(c.Number[:i+1])
			if err != nil {
				return Company{"", ""}, errors.New("Unknown credit card method")
			}
		}
	}

	switch {
	case ccDigits.At(2) == 34 || ccDigits.At(2) == 37:
		return Company{"amex", "American Express"}, nil
	case ccDigits.At(4) == 5610 || (ccDigits.At(6) >= 560221 && ccDigits.At(6) <= 560225):
		return Company{"bankcard", "Bankcard"}, nil
	case ccDigits.At(6) == 604400 || ccDigits.At(6) == 627170 || ccDigits.At(6) == 603522 || ccDigits.At(6) == 589657 || (ccDigits.At(6) >= 604201 && ccDigits.At(6) <= 604219) || (ccDigits.At(6) >= 604300 && ccDigits.At(6) <= 604399):
		return Company{"cabal", "Cabal"}, nil
	case ccDigits.At(2) == 62 || ccDigits.At(2) == 81:
		return Company{"china unionpay", "China UnionPay"}, nil
	case ccDigits.At(3) >= 300 && ccDigits.At(3) <= 305 && ccLen == 15:
		return Company{"diners club carte blanche", "Diners Club Carte Blanche"}, nil
	case ccDigits.At(4) == 2014 || ccDigits.At(4) == 2149:
		return Company{"diners club enroute", "Diners Club enRoute"}, nil
	case ((ccDigits.At(3) >= 300 && ccDigits.At(3) <= 305) || ccDigits.At(3) == 309 || ccDigits.At(2) == 36 || ccDigits.At(2) == 38 || ccDigits.At(2) == 39) && ccLen <= 14:
		return Company{"diners club international", "Diners Club International"}, nil
	case ccDigits.At(4) == 6011 || (ccDigits.At(6) >= 622126 && ccDigits.At(6) <= 622925) || (ccDigits.At(3) >= 644 && ccDigits.At(3) <= 649) || ccDigits.At(2) == 65:
		return Company{"discover", "Discover"}, nil
	// Elo must be checked before interpayment
	case ccDigits.At(4) == 4011 || ccDigits.At(6) == 431274 || ccDigits.At(6) == 438935 ||
		ccDigits.At(6) == 451416 || ccDigits.At(6) == 457393 || ccDigits.At(4) == 4576 ||
		ccDigits.At(6) == 457631 || ccDigits.At(6) == 457632 || ccDigits.At(6) == 504175 ||
		ccDigits.At(6) == 627780 || ccDigits.At(6) == 636297 || ccDigits.At(6) == 636368 ||
		ccDigits.At(6) == 636369 || (ccDigits.At(6) >= 506699 && ccDigits.At(6) <= 506778) ||
		(ccDigits.At(6) >= 509000 && ccDigits.At(6) <= 509999) ||
		(ccDigits.At(6) >= 650031 && ccDigits.At(6) <= 650051) ||
		(ccDigits.At(6) >= 650035 && ccDigits.At(6) <= 650033) ||
		(ccDigits.At(6) >= 650405 && ccDigits.At(6) <= 650439) ||
		(ccDigits.At(6) >= 650485 && ccDigits.At(6) <= 650538) ||
		(ccDigits.At(6) >= 650541 && ccDigits.At(6) <= 650598) ||
		(ccDigits.At(6) >= 650700 && ccDigits.At(6) <= 650718) ||
		(ccDigits.At(6) >= 650720 && ccDigits.At(6) <= 650727) ||
		(ccDigits.At(6) >= 650901 && ccDigits.At(6) <= 650920) ||
		(ccDigits.At(6) >= 651652 && ccDigits.At(6) <= 651679) ||
		(ccDigits.At(6) >= 655000 && ccDigits.At(6) <= 655019) ||
		(ccDigits.At(6) >= 655021 && ccDigits.At(6) <= 655021):
		return Company{"elo", "Elo"}, nil
	case matchesValue(ccDigits.At(6), []int{606282, 637095, 637568, 637599, 637609, 637612}):
		return Company{"hipercard", "Hipercard"}, nil
	case ccDigits.At(3) == 636 && ccLen >= 16 && ccLen <= 19:
		return Company{"interpayment", "InterPayment"}, nil
	case ccDigits.At(3) >= 637 && ccDigits.At(3) <= 639 && ccLen == 16:
		return Company{"instapayment", "InstaPayment"}, nil
	case ccDigits.At(4) >= 3528 && ccDigits.At(4) <= 3589:
		return Company{"jcb", "JCB"}, nil
	case ccDigits.At(6) == 589562:
		return Company{"naranja", "Naranja"}, nil
	case ccDigits.At(4) == 5018 || ccDigits.At(4) == 5020 || ccDigits.At(4) == 5038 || ccDigits.At(4) == 5612 || ccDigits.At(4) == 5893 || ccDigits.At(4) == 6304 || ccDigits.At(4) == 6759 || ccDigits.At(4) == 6761 || ccDigits.At(4) == 6762 || ccDigits.At(4) == 6763 || c.Number[:3] == "0604" || ccDigits.At(4) == 6390:
		return Company{"maestro", "Maestro"}, nil
	case ccDigits.At(4) == 5019:
		return Company{"dankort", "Dankort"}, nil
	case ccDigits.At(2) >= 51 && ccDigits.At(2) <= 55 || ccDigits.At(6) >= 222100 && ccDigits.At(6) <= 272099:
		return Company{"mastercard", "MasterCard"}, nil
	case ccDigits.At(4) == 4026 || ccDigits.At(6) == 417500 || ccDigits.At(4) == 4405 || ccDigits.At(4) == 4508 || ccDigits.At(4) == 4844 || ccDigits.At(4) == 4913 || ccDigits.At(4) == 4917:
		return Company{"visa electron", "Visa Electron"}, nil
	case ccDigits.At(1) == 4:
		return Company{"visa", "Visa"}, nil
	case ccDigits.At(2) == 50:
		return Company{"aura", "Aura"}, nil
	default:
		return Company{"", ""}, errors.New("Unknown credit card method")
	}
}

// Luhn algorithm
// http://en.wikipedia.org/wiki/Luhn_algorithm

// ValidateNumber will check the credit card's number against the Luhn algorithm
func (c *Card) ValidateNumber() bool {
	var sum int
	var alternate bool

	numberLen := len(c.Number)

	if numberLen < 13 || numberLen > 19 {
		return false
	}

	for i := numberLen - 1; i > -1; i-- {
		mod, _ := strconv.Atoi(string(c.Number[i]))
		if alternate {
			mod *= 2
			if mod > 9 {
				mod = (mod % 10) + 1
			}
		}

		alternate = !alternate

		sum += mod
	}

	return sum%10 == 0
}

func matchesValue(number int, numbers []int) bool {
	for _, v := range numbers {
		if v == number {
			return true
		}
	}
	return false
}
