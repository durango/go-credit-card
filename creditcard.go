package creditcard

import (
	"errors"
	"strconv"
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
	timeNow := timeNowCaller()

	if len(c.Year) < 3 {
		year, err = strconv.Atoi(strconv.Itoa(timeNow.UTC().Year())[:2] + c.Year)
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

	if year < timeNowCaller().UTC().Year() {
		return errors.New("Credit card has expired")
	}

	if year == timeNowCaller().UTC().Year() && month < int(timeNowCaller().UTC().Month()) {
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
	case isAmex(ccDigits):
		return Company{"amex", "American Express"}, nil
	case isBankCard(ccDigits):
		return Company{"bankcard", "Bankcard"}, nil
	case isCabal(ccDigits):
		return Company{"cabal", "Cabal"}, nil
	case isUnionPay(ccDigits):
		return Company{"china unionpay", "China UnionPay"}, nil
	case isDinersClubCarteBlanche(ccDigits, ccLen):
		return Company{"diners club carte blanche", "Diners Club Carte Blanche"}, nil
	case isDinersClubEnroute(ccDigits):
		return Company{"diners club enroute", "Diners Club enRoute"}, nil
	case isDinersClubInternational(ccDigits, ccLen):
		return Company{"diners club international", "Diners Club International"}, nil
	case isDiscover(ccDigits):
		return Company{"discover", "Discover"}, nil
	// Elo must be checked before interpayment
	case isElo(ccDigits):
		return Company{"elo", "Elo"}, nil
	case isHipercard(ccDigits):
		return Company{"hipercard", "Hipercard"}, nil
	case isInterpayment(ccDigits, ccLen):
		return Company{"interpayment", "InterPayment"}, nil
	case isInstapayment(ccDigits, ccLen):
		return Company{"instapayment", "InstaPayment"}, nil
	case isJCB(ccDigits):
		return Company{"jcb", "JCB"}, nil
	case isNaranja(ccDigits):
		return Company{"naranja", "Naranja"}, nil
	case isMaestro(c, ccDigits):
		return Company{"maestro", "Maestro"}, nil
	case isDankort(ccDigits):
		return Company{"dankort", "Dankort"}, nil
	case isMasterCard(ccDigits):
		return Company{"mastercard", "MasterCard"}, nil
	case isVisaElectron(ccDigits):
		return Company{"visa electron", "Visa Electron"}, nil
	case isVisa(ccDigits):
		return Company{"visa", "Visa"}, nil
	case isAura(ccDigits):
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

func isInBetween(n, min, max int) bool {
	return n >= min && n <= max
}
