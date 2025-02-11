package main

import (
	"fmt"
	"time"
)

func main() {
	str1 := "hello\nworld\nline\nthhreetoetuhoe\nyes"
	str2 := "bye\nworld\nwidth\nsucks"

	// fmt.Println("str1:\n", str1)
	// fmt.Println()
	// fmt.Println("str2:\n", str2)
	// fmt.Println()

	start := time.Now()
	fmt.Printf("%q\n", JoinVertical(Center,
		str1, str2,
	))
	fmt.Println("Without Grow:", time.Since(start))

	// fmt.Println()

	// fmt.Println(JoinHorizontal(Top,
	// 	str1, str2,
	// ))
}
