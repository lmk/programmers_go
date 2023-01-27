package main

import (
	"fmt"
	"time"
)

func solution(array []int, n int) int {
	count := 0
	for _, a := range array {
		if a == n {
			count++
		}
	}
	return count
}

func report(array []int, height int, desireResult int) {

	input := fmt.Sprintf("Input:[%v, %v]", array, height)

	start := time.Now()

	result := solution(array, height)

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
	report([]int{1, 1, 2, 3, 4, 5}, 1, 2)
	report([]int{0, 2, 3, 4}, 1, 0)
}
