package main

import (
	"fmt"
	"time"
)

func solution(n int, k int) int {
	sum := n * 12000
	if k*10 > n {
		sum += (k - (n / 10)) * 2000
	}

	return sum
}

func report(n int, k int, desireResult int) {

	input := fmt.Sprintf("Input:[%v, %v]", n, k)

	start := time.Now()

	result := solution(n, k)

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
	report(10, 3, 124000)
	report(64, 6, 768000)
}
