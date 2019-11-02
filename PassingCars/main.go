package main

import "fmt"

func main() {
	fmt.Println(Solution([]int{0, 1, 0, 1, 1}))
}

func Solution(A []int) int {
	count, base := 0, 0
	for _, v := range A {
		switch v {
		case 0:
			base++
		case 1:
			count += base
		}
		if count > 1000000000 {
			return -1
		}
	}
	return count
}
