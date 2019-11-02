package main

import "fmt"

func main() {
	fmt.Println(Solution([]int{9, 3, 9, 3, 9, 7, 9}))
}

func Solution(A []int) int {
	res := 0
	for _, n := range A {
		res ^= n
	}
	return res
}
