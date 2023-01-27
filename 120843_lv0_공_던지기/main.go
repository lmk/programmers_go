package main

import (
	"fmt"
	"time"
)

func solution(numbers []int, k int) int {
	return numbers[((k-1)*2)%len(numbers)]
}

func report(numbers []int, k int, desireResult int) {

	input := fmt.Sprintf("Input:[%v, %v]", numbers, k)

	start := time.Now()

	result := solution(numbers, k)

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
	report([]int{1, 2, 3, 4}, 2, 3)
	report([]int{1, 2, 3, 4, 5, 6}, 5, 3)
	report([]int{1, 2, 3}, 3, 2)
}
