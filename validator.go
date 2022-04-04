package creditcard

func isAmex(ccDigits digits) bool {
	return matchesValue(ccDigits.At(2), []int{34, 37})
}

func isBankCard(ccDigits digits) bool {
	return ccDigits.At(4) == 5610 || isInBetween(ccDigits.At(6), 560221, 560225)
}

func isCabal(ccDigits digits) bool {
	atSix := ccDigits.At(6)

	return matchesValue(atSix, []int{604400, 627170, 603522, 589657}) ||
		isInBetween(atSix, 604201, 604219) ||
		isInBetween(atSix, 604300, 604399)
}

func isUnionPay(ccDigits digits) bool {
	return matchesValue(ccDigits.At(2), []int{62, 81})
}

func isDinersClubCarteBlanche(ccDigits digits, ccLen int) bool {
	return isInBetween(ccDigits.At(3), 300, 305) && ccLen == 14
}

func isDinersClubEnroute(ccDigits digits) bool {
	return matchesValue(ccDigits.At(4), []int{2014, 2149})
}

func isDinersClubInternational(ccDigits digits, ccLen int) bool {
	checkThree := isInBetween(ccDigits.At(3), 300, 305) || ccDigits.At(3) == 309
	checkTwoo := matchesValue(ccDigits.At(2), []int{36, 38, 39})

	return (checkThree || checkTwoo) && ccLen <= 14
}

func isDiscover(ccDigits digits) bool {
	return ccDigits.At(4) == 6011 ||
		isInBetween(ccDigits.At(6), 622126, 622925) ||
		isInBetween(ccDigits.At(3), 644, 649) ||
		ccDigits.At(2) == 65
}

func isElo(ccDigits digits) bool {
	atFour := ccDigits.At(4)
	atSix := ccDigits.At(6)

	return matchesValue(atFour, []int{4011, 4576}) ||
		matchesValue(atSix, []int{431274, 438935, 451416, 457393, 457631, 457632, 504175, 627780, 636297, 636368, 636369}) ||
		isInBetween(atSix, 506699, 506778) ||
		isInBetween(atSix, 509000, 509999) ||
		isInBetween(atSix, 650031, 650051) ||
		isInBetween(atSix, 650035, 650033) ||
		isInBetween(atSix, 650405, 650439) ||
		isInBetween(atSix, 650485, 650538) ||
		isInBetween(atSix, 650541, 650598) ||
		isInBetween(atSix, 650700, 650718) ||
		isInBetween(atSix, 650720, 650727) ||
		isInBetween(atSix, 650901, 650920) ||
		isInBetween(atSix, 651652, 651679) ||
		isInBetween(atSix, 655000, 655019) ||
		isInBetween(atSix, 655021, 655021)
}

func isHipercard(ccDigits digits) bool {
	return matchesValue(ccDigits.At(6), []int{606282, 637095, 637568, 637599, 637609, 637612})
}

func isInterpayment(ccDigits digits, ccLen int) bool {
	return ccDigits.At(3) == 636 && isInBetween(ccLen, 16, 19)
}

func isInstapayment(ccDigits digits, ccLen int) bool {
	return isInBetween(ccDigits.At(3), 637, 639) && ccLen == 16
}

func isJCB(ccDigits digits) bool {
	return isInBetween(ccDigits.At(4), 3528, 3589)
}

func isNaranja(ccDigits digits) bool {
	return ccDigits.At(6) == 589562
}

func isMaestro(c *Card, ccDigits digits) bool {
	return matchesValue(ccDigits.At(4), []int{5018, 5020, 5038, 5612, 5893, 6304, 6759, 6761, 6762, 6763, 6390}) ||
		c.Number[:3] == "0604"
}

func isDankort(ccDigits digits) bool {
	return ccDigits.At(4) == 5019
}

func isMasterCard(ccDigits digits) bool {
	return isInBetween(ccDigits.At(2), 51, 55) || isInBetween(ccDigits.At(6), 222100, 272099)
}

func isVisaElectron(ccDigits digits) bool {
	return matchesValue(ccDigits.At(4), []int{4026, 4405, 4508, 4844, 4913, 4917}) || ccDigits.At(6) == 417500
}

func isVisa(ccDigits digits) bool {
	return ccDigits.At(1) == 4
}

func isAura(ccDigits digits) bool {
	return ccDigits.At(2) == 50
}
