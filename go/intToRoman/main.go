package main

import "fmt"

func main() {
	fmt.Println(romanToInt1("MCMXCIV"))
}

var foo = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
var bar = []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

func intToRoman(num int) string {
	var res string
	for i := 0; i < len(foo); i++ {
		if num == 0 {
			return res
		}
		for num >= foo[i] {
			res = res + bar[i]
			num -= foo[i]
		}
	}
	return res
}

func romanToInt(s string) int {
	var res int
	for i, symbol := range bar {
		if s == "" {
			return res
		}
		for {
			if len(s) < len(symbol) {
				break
			}
			if symbol != s[0:len(symbol)] {
				break
			}
			res += foo[i]
			s = s[len(symbol):]
		}
	}
	return res
}

var m = map[uint8]int{
	'M': 1000,
	'D': 500,
	'C': 100,
	'L': 50,
	'X': 10,
	'V': 5,
	'I': 1,
}

func romanToInt1(s string) int {
	var res int
	var prev = 0

	for i := len(s) - 1; i >= 0; i-- {
		val := m[s[i]]
		if val >= prev {
			res += val
		} else {
			res -= val
		}
		prev = val
	}
	return res
}
