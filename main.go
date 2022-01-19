// Package creditcard provides methods for validating credit cards
package creditcard

import "time"

var timeNowCaller = time.Now

func resetMocks() {
	timeNowCaller = time.Now
}
