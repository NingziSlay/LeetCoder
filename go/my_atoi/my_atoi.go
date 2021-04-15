package main

import (
	"fmt"
)

func main() {
	fmt.Println(myAtoi("001-sda0002147483643"))
}

func myAtoi(s string) int {
	var (
		MaxInt   = 1<<31 - 1
		MinInt   = -1 << 31
		started  bool
		negative bool
		result   int
	)

	for _, i := range s {
		switch i {
		case ' ':
			if started {
				break
			}
			continue
		case '+', '-':
			if started {
				break
			}
			negative = i == '-'
			started = true
			continue
		}

		if i < '0' || i > '9' {
			break
		}

		started = true

		e := i - '0'

		if negative {
			if result < MinInt/10 || -(MinInt-result*10) < int(e) {
				return MinInt
			}
			result = result*10 - int(e)
		} else {
			if result > MaxInt/10 || MaxInt-result*10 < int(e) {
				return MaxInt
			}
			result = result*10 + int(e)
		}
	}
	return result
}

func myAtoi1(s string) int {
	var (
		MaxInt   = 1<<31 - 1
		MinInt   = -1 << 31
		started  bool
		negative bool
		result   int
	)

	for _, i := range s {

		// skip blank
		if i == ' ' {
			if !started {
				continue
			} else {
				break
			}
		}
		// negative or not
		if i == '-' {
			if started {
				break
			}
			negative = true
			started = true
			continue
		}
		if i == '+' {
			if started {
				break
			}
			started = true
			continue
		}

		// break if i is not an integer
		if i < '0' || i > '9' {
			break
		}
		started = true
		// calculate result
		e := i - '0'

		if negative {
			if result < MinInt/10 || -(MinInt-result*10) < int(e) {
				return MinInt
			}
			result = result*10 - int(e)
		} else {
			if result > MaxInt/10 || MaxInt-result*10 < int(e) {
				return MaxInt
			}
			result = result*10 + int(e)
		}
	}
	return result
}