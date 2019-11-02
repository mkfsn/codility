package main

import "fmt"

func main() {
	fmt.Println(Solution("CAGCCTA", []int{2, 5, 0}, []int{4, 5, 6}))
	fmt.Println(Solution("G", []int{0}, []int{0}))
}

const MaxN = 100001

func Solution(S string, P []int, Q []int) []int {
	// S: "CAGCCTA"
	positions := map[byte]int{
		'A': MaxN,
		'C': MaxN,
		'G': MaxN,
	}
	matrix := map[byte][]int{
		'A': make([]int, len(S)),
		'C': make([]int, len(S)),
		'G': make([]int, len(S)),
	}
	// matrix
	// A: 1, 1, 6, 6, 6, 6, 6
	// C: 0, 3, 3, 3, 4, N, N
	// G: 2, 2, 3, N, N, N, N
	for i := len(S) - 1; i >= 0; i-- {
		ch := S[i]
		positions[ch] = i
		matrix['A'][i] = positions['A']
		matrix['C'][i] = positions['C']
		matrix['G'][i] = positions['G']
	}

	// case (2, 4)
	// - matrix['A'][2]=6 => 6 > 4 => no A in (2, 4)
	// - matrix['C'][2]=3 => 3 < 4 => has C in (2, 4)
	// - matrix['G'][2]=3 => 3 < 4 => has G in (2, 4)
	//
	// case (5, 5)
	// - matrix['A'][5]=6 => 6 > 5 => no A in (5, 5)
	// - matrix['C'][5]=N => N > 5 => no A in (5, 5)
	// - matrix['G'][5]=N => N > 5 => no A in (5, 5)
	//
	// case (0, 6)
	// - matrix['A'][0]=1 => 1 < 6 => has A in (0, 6)
	//
	res := make([]int, 0, len(P))
	for i := range P {
		l, r := P[i], Q[i]
		switch {
		case matrix['A'][l] <= r:
			res = append(res, 1)
		case matrix['C'][l] <= r:
			res = append(res, 2)
		case matrix['G'][l] <= r:
			res = append(res, 3)
		default:
			res = append(res, 4)
		}
	}
	return res
}
