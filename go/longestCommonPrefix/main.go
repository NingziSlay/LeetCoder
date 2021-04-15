package main

import "fmt"

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flower", "flower", "flower"}))
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	if len(strs) == 1 {
		return strs[0]
	}

	sample := strs[0]

outer:
	for i := len(sample); i >= 0; i-- {
		s := sample[:i]
		for _, str := range strs {
			if len(str) < i {
				continue outer
			}
			if s != str[:i] {
				continue outer
			}
		}
		return s
	}
	return ""
}
