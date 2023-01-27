package main

import (
	"fmt"
	"time"
)

func solution(numbers []int) []int {

	for i := range numbers {
		numbers[i] = numbers[i] * 2
	}

	return numbers
}

func report(numbers []int, desireResult []int) {

	input := fmt.Sprintf("Input:[%v]", numbers)

	start := time.Now()

	result := solution(numbers)

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
	report([]int{1, 2, 3, 4, 5}, []int{2, 4, 6, 8, 10})
	report([]int{1, 2, 100, -99, 1, 2, 3}, []int{2, 4, 200, -198, 2, 4, 6})
}
