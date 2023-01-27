package main

import (
	"fmt"
	"time"
)

func solution(angle int) int {
	if 0 < angle && angle < 90 {
		return 1
	} else if 90 == angle {
		return 2
	} else if 90 < angle && angle < 180 {
		return 3
	} else if 180 == angle {
		return 4
	}

	return 0
}

func report(angle int, desireResult int) {

	input := fmt.Sprintf("Input:[%v]", angle)

	start := time.Now()

	result := solution(angle)

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
	report(70, 1)
	report(91, 3)
	report(180, 4)
}
