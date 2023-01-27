package main

import (
	"fmt"
	"time"
)

func solution(M int, N int) int {
	return M*N - 1
}

func report(m int, n int, desireResult int) {

	input := fmt.Sprintf("Input:[%v, %v]", m, n)

	start := time.Now()

	result := solution(m, n)

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
	report(2, 2, 3)
	report(2, 5, 9)
	report(1, 1, 0)
}
