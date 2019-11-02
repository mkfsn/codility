package main

import "fmt"

func main() {
	fmt.Println(Solution([]int{2, 1, 1, 2, 3, 1}))
}

func Solution(A []int) int {
	m := map[int]int{}
	for _, v := range A {
		m[v]++
	}
	return len(m)
}
