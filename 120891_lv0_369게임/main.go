package main

import (
	"fmt"
	"time"
)

func solution(order int) int {

	c := 0
	for _, r := range fmt.Sprintf("%d", order) {
		if r == '3' || r == '6' || r == '9' {
			c++
		}
	}

	return c
}

func report(order int, desireResult int) {

	input := fmt.Sprintf("Input:[%v]", order)

	start := time.Now()

	result := solution(order)

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
	report(3, 1)
	report(29423, 2)
}
