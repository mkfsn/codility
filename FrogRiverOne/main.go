package main

import "fmt"

func main() {
	fmt.Println(Solution(5, []int{1, 3, 1, 4, 2, 3, 5, 4}))
	fmt.Println(Solution(1, []int{1}))
}

func Solution(X int, A []int) int {
	found := make(map[int]bool)
	total := (1 + X) * X / 2
	for i, v := range A {
		if found[v] {
			continue
		}
		found[v] = true
		total -= v
		if total == 0 {
			return i
		}
	}
	return -1
}
