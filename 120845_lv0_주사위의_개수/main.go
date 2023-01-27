package main

import (
	"fmt"
	"time"
)

func solution(box []int, n int) int {
	return (box[0] / n) * (box[1] / n) * (box[2] / n)
}

func report(box []int, n int, desireResult int) {

	input := fmt.Sprintf("Input:[%v, %v]", box, n)

	start := time.Now()

	result := solution(box, n)

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
	report([]int{1, 1, 1}, 1, 1)
	report([]int{10, 8, 6}, 3, 12)
}
