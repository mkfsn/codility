package main

import "sort"

func main() {

}

func Solution(A []int) int {
	sort.Ints(A)
	for i := 0; i < len(A)-2; i++ {
		if A[i]+A[i+1] > A[i+2] {
			return 1
		}
	}
	return 0
}

// a < b < c
// b + c > a
// a + c > b
// if a + b > c
