package creditcard

import (
	"strconv"
)

// Card holds generic information about the credit card
type Card struct {
	Company Company `json:"company"`
	Cvv     string  `json:"cvv,omitempty"`
	Month   string  `json:"month,omitempty"`
	Number  string  `json:"number,omitempty"`
	Year    string  `json:"year,omitempty"`
}

func NewCard(cvv string, month string, number string, year string) *Card {
	return &Card{Cvv: cvv, Month: month, Number: number, Year: year}
}

// LastFour returns the last four digits of the credit card's number
func (c *Card) LastFour() (string, error) {
	if len(c.Number) < 4 {
		return "", errCardNumberNotLongEnough
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

		return errTestNumbersNotAllowed
	}

	valid := c.ValidateNumber()

	if !valid {
		return errInvalidCardNumber
	}

	return nil
}

// ValidateExpiration validates the credit card's expiration date.
func (c *Card) ValidateExpiration() error {
	var year, month int
	var err error
	timeNow := timeNowCaller()

	if len(c.Year) < 3 {
		year, err = strconv.Atoi(strconv.Itoa(timeNow.UTC().Year())[:2] + c.Year)
		if err != nil {
			return errInvalidYear
		}
	} else {
		year, err = strconv.Atoi(c.Year)
		if err != nil {
			return errInvalidYear
		}
	}

	month, err = strconv.Atoi(c.Month)
	if err != nil {
		return errInvalidMonth
	}

	if month < 1 || 12 < month {
		return errInvalidMonth
	}

	if year < timeNow.UTC().Year() {
		return errCardHasExpired
	}

	if year == timeNow.UTC().Year() && month < int(timeNow.UTC().Month()) {
		return errCardHasExpired
	}

	return nil
}

// ValidateCVV validates the length of the card's CVV value.
func (c *Card) ValidateCVV() error {
	if len(c.Cvv) < 3 || len(c.Cvv) > 4 {
		return errInvalidCVV
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
				return getCompany(""), errUnkownCardMethod
			}
		}
	}

	switch {
	case isAmex(ccDigits):
		return getCompany(CompanyAmericanExpress), nil
	case isBankCard(ccDigits):
		return getCompany(CompanyBankCard), nil
	case isCabal(ccDigits):
		return getCompany(CompanyCabal), nil
	case isUnionPay(ccDigits):
		return getCompany(CompanyChinaUnionPay), nil
	case isDinersClubCarteBlance(ccDigits, ccLen):
		return getCompany(CompanyDinersClubCarteBlance), nil
	case isDinersClubEnroute(ccDigits):
		return getCompany(CompanyDinersClubEnRoute), nil
	case isDinersClubInternational(ccDigits, ccLen):
		return getCompany(CompanyDinersClubInternational), nil
	case isDiscover(ccDigits):
		return getCompany(CompanyDiscover), nil
	// Elo must be checked before interpayment
	case isElo(ccDigits):
		return getCompany(CompanyElo), nil
	case isHipercard(ccDigits):
		return getCompany(CompanyHipercard), nil
	case isInterpayment(ccDigits, ccLen):
		return getCompany(CompanyInterPayment), nil
	case isInstapayment(ccDigits, ccLen):
		return getCompany(CompanyInstaPayment), nil
	case isJCB(ccDigits):
		return getCompany(CompanyJCB), nil
	case isNaranja(ccDigits):
		return getCompany(CompanyNaranja), nil
	case isMaestro(c, ccDigits):
		return getCompany(CompanyMaestro), nil
	case isDankort(ccDigits):
		return getCompany(CompanyDankort), nil
	case isMasterCard(ccDigits):
		return getCompany(CompanyMasterCard), nil
	case isVisaElectron(ccDigits):
		return getCompany(CompanyVisaElectron), nil
	case isVisa(ccDigits):
		return getCompany(CompanyVisa), nil
	case isAura(ccDigits):
		return getCompany(CompanyAura), nil
	default:
		return getCompany(""), errUnkownCardMethod
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
