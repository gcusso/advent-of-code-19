package day04

const numDigits int = 6

var (
	minValue int
	maxValue int
)

func init() {
	minValue = tenPow(numDigits)
	maxValue = tenPow(numDigits+1) - 1
}

func tenPow(e int) int {
	r := 1
	for i := 0; i < e-1; i++ {
		r *= 10
	}
	return r
}

func isValid(v int) bool {
	// It is a six-digit number.
	if v < minValue || v > maxValue {
		return false
	}

	// Two adjacent digits are the same (like 22 in 122345).
	// Going from left to right, the digits never decrease; they only ever increase or stay the same (like 111123 or 135679).
	last := 10
	hasAdjacent := false
	inGroup := false

	for i := 1; i <= numDigits; i++ {
		r := v / tenPow(i) % 10
		if r == last && (hasAdjacent == inGroup) {
			hasAdjacent = true
			if inGroup {
				hasAdjacent = false
			}
			inGroup = true
		} else {
			inGroup = false
		}

		if r > last {
			return false
		}
		last = r
	}

	return hasAdjacent

}

func calcNumPasswords(low, high int) int {

	count := 0

	for cur := low; cur < high; cur++ {
		if isValid(cur) {
			count++
		}
	}

	return count
}
