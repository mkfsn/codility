package main

import (
	"fmt"
)

func main() {
	fmt.Println(solution(32))   // 0
	fmt.Println(solution(9))    // 2
	fmt.Println(solution(529))  // 4
	fmt.Println(solution(20))   // 1
	fmt.Println(solution(15))   // 0
	fmt.Println(solution(1041)) // 5
}

func solution(N int) int {
	var b []int
	for N > 0 {
		b, N = append(b, N%2), N/2
	}

	i := 0
	for ; i < len(b) && b[i] != 1; i++ {
	}

	if i >= len(b) {
		return 0
	}

	max, sum := 0, 0
	for ; i < len(b); i++ {
		switch b[i] {
		case 1:
			if max < sum {
				max = sum
			}
			sum = 0
		case 0:
			sum++
		}
	}
	return max
}
