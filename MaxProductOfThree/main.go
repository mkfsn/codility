package main

import "sort"

func main() {

}

func Solution(A []int) int {
	sort.Ints(A)
	L := len(A)
	x := A[0] * A[1] * A[L-1]
	y := A[L-3] * A[L-2] * A[L-1]
	if x > y {
		return x
	}
	return y
}
