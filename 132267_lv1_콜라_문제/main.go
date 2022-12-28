package main

import (
	"fmt"
	"time"
)

func solution(a int, b int, n int) int {
	answer := 0
	mod := 0
	for n >= a {
		mod = n % a
		n = n / a * b
		answer += n
		n += mod

	}
	return answer
}

func report(a int, b int, n int, desireResult int) {

	input := fmt.Sprintf("Input:[%v %v %v]", a, b, n)

	start := time.Now()

	result := solution(a, b, n)

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
	report(2, 1, 20, 19)
	report(3, 1, 20, 9)
}
