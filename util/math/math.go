// This package provides some basic math utilies.
package math

// Returns the max of the two given ints.
func Max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// Returns the min of the two given ints.
func Min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

// Returns the absolute value of the given int.
func Abs(a int) int {
	if a >= 0 {
		return a
	} else {
		return a * -1
	}
}

// Returns -1 if the given int is less than 0,
//	    1 if the given int is greater than 0,
//	    0 otherwise.
func Signum(a int) int {
	if a < 0 {
		return -1
	} else if a > 0 {
		return 1
	} else {
		return 0
	}
}

// Returns the absolute value of the difference
// between ints a and b.
func AbsDifference(a int, b int) int {
	d1 := a - b
	d2 := b - a
	return Max(d1, d2)
}
