package main

import "fmt"

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

type StringState struct {
	flag   []bool
	ccount int // count of distinct characters
}

// build the state of a string
func Build(str string, stat *StringState) {
	stat.ccount = 0
	for _, c := range str {
		key := c - 'a'
		if !stat.flag[key] {
			stat.flag[key] = true
			stat.ccount++
		}
	}
}

// test if a can be made from b
// if a is made from b, return true
// otherwise, return false
func (x *StringState) Test(a *StringState) bool {
	if a.ccount > x.ccount {
		return false
	}
	var count int
	for i := 0; i < 10; i++ {
		if a.flag[i] && x.flag[i] {
			count++
		}
	}
	return count >= a.ccount
}

// union a to x
func (x *StringState) Union(a *StringState) {
	for i := 0; i < 10; i++ {
		if a.flag[i] && !x.flag[i] {
			x.flag[i] = true
			x.ccount++
		}
	}
}

func Solution(S []string, K int) int {
	var flags = make([]StringState, len(S))
	for i, str := range S {
		// each string is made of lowercase letters from 'a' to 'j'
		flags[i] = StringState{make([]bool, 10), 0}
		Build(str, &flags[i])
	}

	var max_result int
	for i := 0; i < len(S); i++ {
		if flags[i].ccount > K {
			// the string cannot be made from K distinct characters
			continue
		}
		// suppose the current string is chosen, build the target set
		var next = StringState{make([]bool, 10), 0}
		var target = StringState{make([]bool, 10), 0}
		for t := i; t < len(S); t++ {
			next.Union(&flags[t])
			if next.ccount > K {
				break
			}
			copy(target.flag, next.flag)
			target.ccount = next.ccount
			if target.ccount == K {
				break
			}
		}
		// test result of the target set
		var j int
		var result int
		for j = 0; j < len(S); j++ {
			if i == j {
				result++
				continue
			}
			ok := target.Test(&flags[j])
			if ok {
				result++
			}
		}
		if result > max_result {
			max_result = result
		}
	}
	return max_result
}

func main() {
	if r := Solution([]string{"abc", "abb", "cb", "a", "bbb"}, 2); r != 3 {
		fmt.Println("Test failed")
	}
	if r := Solution([]string{"adf", "jjbh", "jcgj", "eijj", "adf"}, 3); r != 2 {
		fmt.Println("Test failed")
	}
	if r := Solution([]string{"abcd", "efgh"}, 3); r != 0 {
		fmt.Println("Test failed")
	}
	if r := Solution([]string{"bc", "edf", "fde", "dge", "abcd"}, 4); r != 3 {
		fmt.Println("Test failed")
	}
}
