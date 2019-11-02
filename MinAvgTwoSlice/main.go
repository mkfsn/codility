package main

import "fmt"

func main() {
	fmt.Println(Solution([]int{4, 2, 2, 5, 1, 5, 8})) // 1
}

func Solution(A []int) int {
	min := float64(A[0]+A[1]) / 2.0
	pos := 0
	i := 0
	for ; i < len(A)-2; i++ {
		if x := float64(A[i]+A[i+1]) / 2.0; x < min {
			min = x
			pos = i
		}
		if x := float64(A[i]+A[i+1]+A[i+2]) / 3.0; x < min {
			min = x
			pos = i
		}
	}
	if x := float64(A[i]+A[i+1]) / 2.0; x < min {
		pos = i
	}
	return pos
}
