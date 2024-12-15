package common

func AbsInt(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func NumDigits(n int) int {
	count := 0
	for n != 0 {
		n /= 10
		count++
	}
	return count
}
