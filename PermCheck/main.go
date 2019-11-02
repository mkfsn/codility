package main

import "fmt"

func main() {
	fmt.Println(Solution([]int{4, 1, 3, 2}))
	fmt.Println(Solution([]int{4, 1, 3}))
}

func Solution(A []int) int {
	table := make(map[int]int)
	sum := 0
	for i, v := range A {
		if table[v] > 0 {
			return 0
		}
		table[v]++
		sum += v - (i + 1)
	}
	if sum == 0 {
		return 1
	}
	return 0
}
