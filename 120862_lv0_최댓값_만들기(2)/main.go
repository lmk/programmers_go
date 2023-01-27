package main

import (
	"fmt"
	"sort"
	"time"
)

func solution(numbers []int) int {

	l := len(numbers)

	sort.Slice(numbers, func(i, j int) bool {
		return numbers[i] > numbers[j]
	})

	if numbers[0]*numbers[1] < numbers[l-2]*numbers[l-1] {
		return numbers[l-2] * numbers[l-1]
	}

	return numbers[0] * numbers[1]
}

func report(numbers []int, desireResult int) {

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
	report([]int{1, 2, -3, 4, -5}, 15)
	report([]int{0, -31, 24, 10, 1, 9}, 240)
	report([]int{10, 20, 30, 5, 5, 20, 5}, 600)
}
