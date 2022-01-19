package creditcard

func isAmex(ccDigits digits) bool {
	return ccDigits.At(2) == 34 || ccDigits.At(2) == 37
}

func isBankCard(ccDigits digits) bool {
	return ccDigits.At(4) == 5610 || (ccDigits.At(6) >= 560221 && ccDigits.At(6) <= 560225)
}

func isCabal(ccDigits digits) bool {
	return ccDigits.At(6) == 604400 || ccDigits.At(6) == 627170 || ccDigits.At(6) == 603522 || ccDigits.At(6) == 589657 || (ccDigits.At(6) >= 604201 && ccDigits.At(6) <= 604219) || (ccDigits.At(6) >= 604300 && ccDigits.At(6) <= 604399)
}

func isUnionPay(ccDigits digits) bool {
	return ccDigits.At(2) == 62 || ccDigits.At(2) == 81
}

func isDinersClubCarteBlance(ccDigits digits, ccLen int) bool {
	return ccDigits.At(3) >= 300 && ccDigits.At(3) <= 305 && ccLen == 15
}

func isDinersClubEnroute(ccDigits digits) bool {
	return ccDigits.At(4) == 2014 || ccDigits.At(4) == 2149
}

func isDinersClubInternational(ccDigits digits, ccLen int) bool {
	return ((ccDigits.At(3) >= 300 && ccDigits.At(3) <= 305) || ccDigits.At(3) == 309 || ccDigits.At(2) == 36 || ccDigits.At(2) == 38 || ccDigits.At(2) == 39) && ccLen <= 14
}

func isDiscover(ccDigits digits) bool {
	return ccDigits.At(4) == 6011 || (ccDigits.At(6) >= 622126 && ccDigits.At(6) <= 622925) || (ccDigits.At(3) >= 644 && ccDigits.At(3) <= 649) || ccDigits.At(2) == 65
}

func isElo(ccDigits digits) bool {
	return ccDigits.At(4) == 4011 || ccDigits.At(6) == 431274 || ccDigits.At(6) == 438935 ||
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
		(ccDigits.At(6) >= 655021 && ccDigits.At(6) <= 655021)
}

func isHipercard(ccDigits digits) bool {
	return matchesValue(ccDigits.At(6), []int{606282, 637095, 637568, 637599, 637609, 637612})
}

func isInterpayment(ccDigits digits, ccLen int) bool {
	return ccDigits.At(3) == 636 && ccLen >= 16 && ccLen <= 19
}

func isInstapayment(ccDigits digits, ccLen int) bool {
	return ccDigits.At(3) >= 637 && ccDigits.At(3) <= 639 && ccLen == 16
}

func isJCB(ccDigits digits) bool {
	return ccDigits.At(4) >= 3528 && ccDigits.At(4) <= 3589
}

func isNaranja(ccDigits digits) bool {
	return ccDigits.At(6) == 589562
}

func isMaestro(c *Card, ccDigits digits) bool {
	return ccDigits.At(4) == 5018 || ccDigits.At(4) == 5020 || ccDigits.At(4) == 5038 || ccDigits.At(4) == 5612 || ccDigits.At(4) == 5893 || ccDigits.At(4) == 6304 || ccDigits.At(4) == 6759 || ccDigits.At(4) == 6761 || ccDigits.At(4) == 6762 || ccDigits.At(4) == 6763 || c.Number[:3] == "0604" || ccDigits.At(4) == 6390
}

func isDankort(ccDigits digits) bool {
	return ccDigits.At(4) == 5019
}

func isMasterCard(ccDigits digits) bool {
	return ccDigits.At(2) >= 51 && ccDigits.At(2) <= 55 || ccDigits.At(6) >= 222100 && ccDigits.At(6) <= 272099
}

func isVisaElectron(ccDigits digits) bool {
	return ccDigits.At(4) == 4026 || ccDigits.At(6) == 417500 || ccDigits.At(4) == 4405 || ccDigits.At(4) == 4508 || ccDigits.At(4) == 4844 || ccDigits.At(4) == 4913 || ccDigits.At(4) == 4917
}

func isVisa(ccDigits digits) bool {
	return ccDigits.At(1) == 4
}

func isAura(ccDigits digits) bool {
	return ccDigits.At(2) == 50
}
