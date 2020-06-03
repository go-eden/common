package emath

func DivCeil(dividend, divisor int) (result int) {
	if divisor == 0 {
		panic("divison can't be 0")
	}
	result = dividend / divisor
	if dividend%divisor != 0 {
		result++
	}
	return
}

func DivRound(dividend, divisor int) (result int) {
	if divisor == 0 {
		panic("divison can't be 0")
	}
	result = dividend / divisor
	if dividend%divisor >= divisor/2 {
		result++
	}
	return
}
