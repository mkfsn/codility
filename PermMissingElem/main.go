package main

import "fmt"

func main() {
	fmt.Println(Solution([]int{2, 3, 1, 5}))
}

func Solution(A []int) int {
	n := len(A) + 1
	sum := (1 + n) * n / 2
	for _, v := range A {
		sum -= v
	}
	return sum
}
