package main

import (
	"fmt"
	"time"
)

func solution(hp int) int {

	count := hp / 5
	hp = hp % 5

	count += hp / 3

	count += hp % 3

	return count
}

func report(hp int, desireResult int) {

	input := fmt.Sprintf("Input:[%v]", hp)

	start := time.Now()

	result := solution(hp)

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
	report(23, 5)
	report(24, 6)
	report(999, 201)
}
