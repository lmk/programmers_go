package main

import (
	"fmt"
	"time"
)

func solution(numbers []int, num1 int, num2 int) []int {
	return numbers[num1 : num2+1]
}

func report(numbers []int, num1 int, num2 int, desireResult []int) {

	input := fmt.Sprintf("Input:[%v]", numbers, num1, num2)

	start := time.Now()

	result := solution(numbers, num1, num2)

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
	report([]int{1, 2, 3, 4, 5}, 1, 3, []int{2, 3, 4})
	report([]int{1, 3, 5}, 1, 2, []int{3, 5})
}
