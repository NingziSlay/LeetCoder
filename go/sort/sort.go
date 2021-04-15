package main

import "fmt"

func main() {
	//       16 8 4 2 1
	// 	0 0 0 0 0 0 1 1 = 3
	//  1 1 1 1 1 1 0 0  取反
	//  0 0 0 0 1 1 1 1
	//  0 0 0 0 1 1 0 0
	var a, b = 15, 3
	fmt.Println(a &^ b)
}

// 归并排序
func Sort(l []int) []int {
	length := len(l)
	if length == 1 {
		return l
	}
	half := length / 2
	var left, right = make([]int, half), make([]int, length-half)

	copy(left, l[:half])
	copy(right, l[half:])
	return merge(l, Sort(left), Sort(right))
}

func merge(m, l, r []int) []int {
	x, y := 0, 0
	lMax, rMax := len(l), len(r)
	for i := 0; i < len(m); i++ {
		if x >= lMax {
			copy(m[i:], r[y:])
			return m
		}
		if y >= rMax {
			copy(m[i:], l[x:])
			return m
		}
		if l[x] > r[y] {
			m[i] = r[y]
			y += 1
		} else {
			m[i] = l[x]
			x += 1
		}
	}
	return m
}
