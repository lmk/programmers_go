package main

import (
	"fmt"
	"time"
)

func solution(numbers []int) float64 {
	sum := 0
	for _, n := range numbers {
		sum += n
	}
	return float64(sum) / float64(len(numbers))
}

func report(numbers []int, desireResult float64) {

	input := fmt.Sprintf("Input:[%v]", numbers)

	start := time.Now()

	result := solution(numbers)

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
	report([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5.5)
	report([]int{89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99}, 94.0)
}
