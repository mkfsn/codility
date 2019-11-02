package main

import (
	"fmt"
)

func main() {
	fmt.Println(Solution([]int{3, 1, 2, 4, 3}))
	fmt.Println(Solution([]int{3, 1}))
}

func Solution(A []int) int {
	r, l := A[0], 0
	for i := 1; i < len(A); i++ {
		l += A[i]
	}

	min := abs(r - l)
	for i := 1; i < len(A)-1; i++ {
		r += A[i]
		l -= A[i]
		if n := abs(r - l); n < min {
			min = n
		}
	}
	return min
}

func abs(n int) int {
	if n >= 0 {
		return n
	}
	return -n
}
