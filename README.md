# go-credit-card [![Build Status](https://travis-ci.org/durango/go-credit-card.svg?branch=master)](https://travis-ci.org/durango/go-credit-card) [![Coverage Status](https://coveralls.io/repos/durango/go-credit-card/badge.svg?branch=master&service=github)](https://coveralls.io/github/durango/go-credit-card?branch=master) [![GoDoc](https://godoc.org/github.com/durango/go-credit-card?status.svg)](https://godoc.org/github.com/durango/go-credit-card)

Basic credit card validation using the [Luhn algorithm](http://en.wikipedia.org/wiki/Luhn_algorithm)

Currently identifies the following credit card companies:
* American Express
* Bankcard
* China Unionpay
* Dankort
* Diners Club Carte Blanche
* Diners Club enRoute
* Diners Club International
* Discover
* InterPayment
* InstaPayment
* JCB
* Maestro
* MasterCard
* Visa
* Visa Electron

## Installation

```bash
go get github.com/durango/go-credit-card
```

## Quick Start

```go
// Initialize a new card:
card := creditcard.Card{Number: "4242424242424242", Cvv: "11111", Month: "02", Year: "2016"}

// Retrieve the card's method (which credit card company this card belongs to)
err := card.Method() // card.Company({Short: "visa", Long: "Visa"})

// Display last four digits
lastFour, err := card.LastFour() // 4242

// Validate the card's number (without capturing)
err := card.Validate() // will return an error due to not allowing test cards

err := card.Validate(true) // this will work though
```
