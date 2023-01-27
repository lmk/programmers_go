package main

import (
	"fmt"
	"time"
)

func solution(n int) int {
	sum := 0
	for i := 2; i <= n; i += 2 {
		sum += i
	}
	return sum
}

func report(angle int, desireResult int) {

	input := fmt.Sprintf("Input:[%v]", angle)

	start := time.Now()

	result := solution(angle)

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
	report(10, 30)
	report(4, 6)
}
