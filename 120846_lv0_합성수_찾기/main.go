package main

import (
	"fmt"
	"time"
)

func isM(n int) bool {
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			return true
		}
	}

	return false
}

func solution(n int) int {

	if n == 1 {
		return 0
	}

	count := 0
	for i := 2; i <= n; i++ {
		if isM(i) {
			count++
		}
	}

	return count
}

func report(num int, desireResult int) {

	input := fmt.Sprintf("Input:[%v]", num)

	start := time.Now()

	result := solution(num)

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
	report(10, 5)
	report(15, 8)
}
