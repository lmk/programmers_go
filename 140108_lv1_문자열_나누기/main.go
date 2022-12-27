package main

import (
	"fmt"
	"time"
)

func mySplit(s string, c int) int {
	if len(s) == 0 {
		return c
	}

	x := s[0]
	xc := 1
	nc := 0

	for i := 1; i < len(s); i++ {
		if x == s[i] {
			xc++
		} else {
			nc++
		}

		if xc == nc {
			c = mySplit(s[i+1:], c)
			break
		}
	}

	return c + 1
}

func solution(s string) int {

	return mySplit(s, 0)
}

func report(s string, desireResult int) {

	input := fmt.Sprintf("Input:[%v]", s)

	start := time.Now()

	result := solution(s)

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
	report("banana", 3)
	report("abracadabra", 6)
	report("aaabbaccccabba", 3)
}
