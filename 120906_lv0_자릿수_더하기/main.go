package main

import (
	"fmt"
	"time"
)

func solution(n int) int {
	sum := 0
	s := fmt.Sprint(n)

	for _, r := range s {
		sum += int(r - '0')
	}

	return sum
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
	report(1234, 10)
	report(930211, 16)
}
