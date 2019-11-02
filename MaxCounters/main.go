package main

import "fmt"

func main() {
	fmt.Println(Solution(5, []int{3, 4, 4, 6, 1, 4, 4}))
}

func Solution(N int, A []int) []int {
	res := make([]int, N)
	max, base := 0, 0
	for _, x := range A {
		switch {
		case x == N+1:
			base = max

		case x >= 1 && x <= N:
			i := x - 1
			v := res[i]
			if v < base {
				v = base + 1
			} else {
				v = v + 1
			}
			res[i] = v

			if v > max {
				max = v
			}
		}
	}

	for i, v := range res {
		if v < base {
			res[i] = base
		}
	}

	return res
}
