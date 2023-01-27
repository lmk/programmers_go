package main

import (
	"fmt"
	"time"
)

func solution(numbers []int, direction string) []int {

	l := len(numbers)
	if direction == "right" {
		return append(numbers[l-1:], numbers[:l-1]...)
	}

	return append(numbers[1:], numbers[0:1]...)
}

func report(numbers []int, direction string, desireResult []int) {

	input := fmt.Sprintf("Input:[%v, %v]", numbers, direction)

	start := time.Now()

	result := solution(numbers, direction)

	duration := time.Since(start)

	isSame := true
	if len(result) != len(desireResult) {
		isSame = false
	} else {
		for i, _ := range result {
			if result[i] != desireResult[i] {
				isSame = false
				break
			}
		}
	}

	if isSame {
		fmt.Printf("[O] %s result:[%v] duration:%v\n", input, result, duration)
	} else {
		fmt.Printf("[X] %s result:[%v]->[%v] duration:%v\n", input, result, desireResult, duration)
	}
}

func main() {
	report([]int{1, 2, 3}, "right", []int{3, 1, 2})
	report([]int{4, 455, 6, 4, -1, 45, 6}, "left", []int{455, 6, 4, -1, 45, 6, 4})
}
