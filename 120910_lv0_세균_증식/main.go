package main

import (
	"fmt"
	"math"
	"time"
)

func solution(n int, t int) int {
	return int(math.Pow(2, float64(t))) * n
}

func report(n int, t int, desireResult int) {

	input := fmt.Sprintf("Input:[%v, %v]", n, t)

	start := time.Now()

	result := solution(n, t)

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
	report(2, 10, 2048)
	report(7, 15, 229376)
}
