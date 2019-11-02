package main

import (
	"fmt"
)

func main() {
	fmt.Println(Solution(6, 11, 2))
	fmt.Println(Solution(11, 345, 17))
	fmt.Println(Solution(0, 1, 11))
	fmt.Println(Solution(10, 10, 5))
	fmt.Println(Solution(10, 10, 7))
	fmt.Println(Solution(10, 10, 20))
}

func Solution(A int, B int, K int) int {
	x, y := A%K, B%K
	if x != 0 {
		A = A + (K - x)
	}
	if y != 0 {
		B = B - y
	}
	return (B-A)/K + 1
}
