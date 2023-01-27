package main

import (
	"fmt"
	"time"
)

func solution(n int) int {

	for i := 1; i <= n*6; i++ {
		if (6*i)%n == 0 {
			return i
		}
	}

	return n
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
	report(6, 1)
	report(10, 5)
	report(4, 2)
}
