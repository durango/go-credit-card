package creditcard

type digits [6]int

// At returns the digits from the start to the given length
func (d *digits) At(i int) int {
	return d[i-1]
}
