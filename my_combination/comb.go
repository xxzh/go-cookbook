package main

import "fmt"

func Comb(s string, k int) []string {
	var result []string
	if k == 0 {
		return result
	}
	if k == 1 {
		for _, c := range s {
			result = append(result, string(c))
		}
		return result
	}

	if len(s) == k {
		result = append(result, s)
		return result
	}

	for i, c := range s {
		var next = Comb(s[i+1:], k-1)
		for _, n := range next {
			result = append(result, string(c)+n)
		}
	}

	return result
}

//
// Comb2 generates all combinations of k elements from a set of n elements.
// The elements are represented by integers from 0 to n-1.
// The result is a slice of slices of integers.
// For example, Comb2(3, 2) returns [[0 1] [0 2] [1 2]].
//
func Comb2(n, k int) [][]int {
	var result [][]int
	var comb []int = make([]int, 0, k)
	var backtrack func(int)
	backtrack = func(start int) {
		if len(comb) == k {
			var c = make([]int, k)
			copy(c, comb)
			result = append(result, c)
			return
		}
		for i := start; i < n; i++ {
			comb = append(comb, i)
			backtrack(i + 1)
			comb = comb[:len(comb)-1]
		}
	}
	backtrack(0)
	return result
}

func main() {
	if r := Comb("abc", 1); len(r) != 3 {
		fmt.Println(r)
	}
	if r := Comb("abc", 2); len(r) != 3 {
		fmt.Println(r)
	}
	if r := Comb("abc", 3); len(r) != 1 {
		fmt.Println(r)
	}
	if r := Comb("abcd", 3); len(r) != 4 {
		fmt.Println(r)
	}
	if r := Comb("abcd", 2); len(r) != 6 {
		fmt.Println(r)
	}
	if r := Comb2(4, 2); len(r) != 4 {
		fmt.Println(r)
	}
}
