package main

import (
	"fmt"
	"time"
)

func solution(n int) int {

	count := 0
	for i := 1; i <= n/2; i++ {
		if n/i*i == n {
			count++
		}
	}

	return count + 1
}

func report(n int, desireResult int) {

	input := fmt.Sprintf("Input:[%v]", n)

	start := time.Now()

	result := solution(n)

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
	report(20, 6)
	report(100, 9)
}
