package main

import (
	"fmt"
	"sort"
	"time"
)

func solution(s string) string {

	m := map[rune]int{}
	for _, r := range s {
		m[r] = m[r] + 1
	}

	rs := []rune{}
	for k, v := range m {
		if v == 1 {
			rs = append(rs, k)
		}
	}

	sort.Slice(rs, func(i, j int) bool {
		return rs[i] < rs[j]
	})

	return string(rs)
}

func report(my_string string, desireResult string) {

	input := fmt.Sprintf("Input:[%v]", my_string)

	start := time.Now()

	result := solution(my_string)

	duration := time.Since(start)

	isSame := true
	if result != desireResult {
		isSame = false
	}

	if isSame {
		fmt.Printf("[O] %s result:[%v] duration:%v\n", input, result, duration)
	} else {
		fmt.Printf("[X] %s result:[%v]->[%v] duration:%v\n", input, result, desireResult, duration)
	}
}

func main() {
	report("abcabcadc", "d")
	report("abdc", "abcd")
	report("hello", "eho")
}
