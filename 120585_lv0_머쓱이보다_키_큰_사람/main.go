package main

import (
	"fmt"
	"time"
)

func solution(array []int, height int) int {
	count := 0
	for _, t := range array {
		if t > height {
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
	report([]int{149, 180, 192, 170}, 167, 3)
	report([]int{180, 120, 140}, 190, 0)
}
