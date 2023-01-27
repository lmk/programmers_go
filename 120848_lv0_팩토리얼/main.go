package main

import (
	"fmt"
	"time"
)

func solution(n int) int {

	i := 1
	c := 1

	for i <= n {
		c++
		i *= c
	}

	return c - 1
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
	report(3628800, 10)
	report(7, 3)
}
