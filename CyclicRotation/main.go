package main

import "fmt"

func main() {
	fmt.Println(Solution([]int{3, 8, 9, 7, 6}, 3))
	fmt.Println(Solution([]int{0, 0, 0}, 1))
	fmt.Println(Solution([]int{5, -1000}, 1))
}

func Solution(A []int, K int) []int {
	if len(A) == 0 {
		return A
	}
	K %= len(A)
	if K == 0 {
		return A
	}
	return append(A[len(A)-K:], A[0:len(A)-K]...)
}
