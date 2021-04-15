package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(reverse(7463847413))
}

func reverse(x int) int {
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
		if r > math.MaxInt32 {
			return 0
		}
		x /= 10
	}
	if negative {
		return -r
	}
	return r
}
