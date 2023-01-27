package main

import (
	"fmt"
	"time"
)

func solution(num int) int {
	if num%7 == 0 {
		return num / 7
	}

	return num/7 + 1
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
	report(7, 1)
	report(1, 1)
	report(15, 3)
}
