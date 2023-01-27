package main

import (
	"fmt"
	"time"
)

func solution(age int) int {
	return 2022 - age + 1
}

func report(age int, desireResult int) {

	input := fmt.Sprintf("Input:[%v]", age)

	start := time.Now()

	result := solution(age)

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
	report(40, 1983)
	report(23, 2000)
}
