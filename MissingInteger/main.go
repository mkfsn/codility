package main

import "fmt"

func main() {
	fmt.Println(Solution([]int{1, 3, 6, 4, 1, 2})) // 5
	fmt.Println(Solution([]int{-1, -3}))           // 1
	fmt.Println(Solution([]int{1, 2, 3}))          // 4
}

func Solution(A []int) int {
	found := make(map[int]bool)
	min := 1
	for _, v := range A {
		found[v] = true
		i := min
		for found[i] {
			i++
		}
		min = i
	}
	return min
}
