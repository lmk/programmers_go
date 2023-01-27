package main

import (
	"fmt"
	"time"
)

func solution(n int) []int {

	answer := []int{}

	for i := 1; i <= n; i += 2 {
		answer = append(answer, i)
	}

	return answer
}

func report(n int, desireResult []int) {

	input := fmt.Sprintf("Input:[%v]", n)

	start := time.Now()

	result := solution(n)

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
	report(10, []int{1, 3, 5, 7, 9})
	report(15, []int{1, 3, 5, 7, 9, 11, 13, 15})
}
