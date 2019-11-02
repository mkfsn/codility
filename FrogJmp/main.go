package main

import "fmt"

func main() {
	fmt.Println(Solution(10, 85, 30))
}

func Solution(X int, Y int, D int) int {
	diff := Y - X
	n, r := diff/D, diff%D
	if r == 0 {
		return n
	}
	return n + 1
}
