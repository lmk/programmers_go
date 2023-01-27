package main

import (
	"fmt"
	"time"
)

func solution(num1 int, num2 int) int {
	return int(float64(num1) / float64(num2) * 1000.0)
}

func report(num1 int, num2 int, desireResult int) {

	input := fmt.Sprintf("Input:[%v, %v]", num1, num2)

	start := time.Now()

	result := solution(num1, num2)

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
	report(3, 2, 1500)
	report(7, 3, 2333)
	report(1, 16, 62)
}
