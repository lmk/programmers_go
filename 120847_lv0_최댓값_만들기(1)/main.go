package main

import (
	"fmt"
	"time"
)

func solution(sides []int) int {
	if sides[0] > sides[1] {
		sides[0], sides[1] = sides[1], sides[0]
	}
	if sides[1] > sides[2] {
		sides[1], sides[2] = sides[2], sides[1]
	}

	if sides[0]+sides[1] > sides[2] {
		return 1
	}

	return 2
}

func report(sides []int, desireResult int) {

	input := fmt.Sprintf("Input:[%v]", sides)

	start := time.Now()

	result := solution(sides)

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
	report([]int{1, 2, 3}, 2)
	report([]int{3, 6, 2}, 2)
	report([]int{199, 72, 222}, 1)
}
