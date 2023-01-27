package main

import (
	"fmt"
	"math"
	"time"
)

func solution(n int) int {

	a := int(math.Sqrt(float64(n)))

	if a*a == n {
		return 1
	}

	return 2
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
	report(144, 1)
	report(976, 2)
}
