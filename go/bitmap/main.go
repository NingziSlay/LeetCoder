package main

import "fmt"

func main() {
	b := NewBitmap(1)
	b.Add(1)
	fmt.Println(len(b.bitmap))
	b.Add(32)
	fmt.Println(len(b.bitmap))
	fmt.Println(b)
}

// 0000000000000000000000010000001000000000000000000000000000000001