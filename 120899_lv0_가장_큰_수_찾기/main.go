package main

import (
	"fmt"
	"time"
)

func solution(array []int) []int {
	answer := []int{array[0], 0}

	for i, n := range array {
		if answer[0] < n {
			answer[0] = n
			answer[1] = i
		}
	}

	return answer
}

func report(array []int, desireResult []int) {

	input := fmt.Sprintf("Input:[%v]", array)

	start := time.Now()

	result := solution(array)

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
	report([]int{1, 8, 3}, []int{8, 1})
	report([]int{9, 10, 11, 8}, []int{11, 2})
}
