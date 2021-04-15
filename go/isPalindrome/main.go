package main

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}

	return x == reverse(x)
}

func reverse(x int) int {
	MaxInt := 1<<31 - 1

	if x < 10 && x > -10 {
		return x
	}

	var negative = false
	if x < 0 {
		negative = true
		x = -x
	}

	var r = 0

	for x != 0 {
		r = r*10 + x%10
		if r > MaxInt {
			return 0
		}
		x /= 10
	}
	if negative {
		return -r
	}
	return r
}
