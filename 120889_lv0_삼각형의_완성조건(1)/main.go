package main

import (
	"fmt"
	"time"
)

func solution(s1 []string, s2 []string) int {

	if len(s1) > len(s2) {
		s1, s2 = s2, s1
	}

	count := 0
	for _, o := range s1 {
		for _, t := range s2 {
			if o == t {
				count++
				continue
			}
		}
	}

	return count
}

func report(s1 []string, s2 []string, desireResult int) {

	input := fmt.Sprintf("Input:[%v, %v]", s1, s2)

	start := time.Now()

	result := solution(s1, s2)

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
	report([]string{"a", "b", "c"}, []string{"com", "b", "d", "p", "c"}, 2)
	report([]string{"n", "omg"}, []string{"m", "dot"}, 0)
}
