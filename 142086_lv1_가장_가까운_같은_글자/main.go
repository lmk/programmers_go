package main

import (
	"fmt"
	"time"
)

func solution(s string) []int {
	m := map[rune]int{}
	answer := make([]int, len(s))

	for i, c := range s {
		if _, ok := m[c]; ok {
			answer[i] = i - m[c]
		} else {
			answer[i] = -1
		}

		m[c] = i
	}

	return answer
}

func report(s string, desireResult []int) {

	input := fmt.Sprintf("Input:[%v]", s)

	start := time.Now()

	result := solution(s)

	duration := time.Since(start)

	isSame := true
	for i := range desireResult {
		if result[i] != desireResult[i] {
			isSame = false
			break
		}
	}

	if isSame {
		fmt.Printf("[O] %s result:[%v] duration:%v\n", input, result, duration)
	} else {
		fmt.Printf("[X] %s result:[%v]->[%v] duration:%v\n", input, result, desireResult, duration)
	}
}

func main() {
	report("banana", []int{-1, -1, -1, 2, 2, 2})
	report("foobar", []int{-1, -1, 1, -1, -1, -1})
}
