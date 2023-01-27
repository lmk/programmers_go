package main

import (
	"fmt"
	"time"
)

func solution(n int) []int {
	answer := []int{}

	for i := 1; i <= n/2; i++ {
		if n%i == 0 {
			answer = append(answer, i)
		}
	}

	return append(answer, n)
}

func report(num int, desireResult []int) {

	input := fmt.Sprintf("Input:[%v]", num)

	start := time.Now()

	result := solution(num)

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
	report(24, []int{1, 2, 3, 4, 6, 8, 12, 24})
	report(29, []int{1, 29})
}
