package main

import (
	"fmt"
	"time"
)

func solution(message string) int {
	return len(message) * 2
}

func report(message string, desireResult int) {

	input := fmt.Sprintf("Input:[%v]", message)

	start := time.Now()

	result := solution(message)

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
	report("happy birthday!", 30)
	report("I love you~", 22)
}
