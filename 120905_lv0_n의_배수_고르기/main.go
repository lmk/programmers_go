package main

import (
	"fmt"
	"time"
)

func solution(n int, numlist []int) []int {

	answer := []int{}
	for _, l := range numlist {
		if l/n*n == l {
			answer = append(answer, l)
		}
	}

	return answer
}

func report(n int, numlist []int, desireResult []int) {

	input := fmt.Sprintf("Input:[%v, %v]", n, numlist)

	start := time.Now()

	result := solution(n, numlist)

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
	report(3, []int{4, 5, 6, 7, 8, 9, 10, 11, 12}, []int{6, 9, 12})
	report(5, []int{1, 9, 3, 10, 13, 5}, []int{10, 5})
	report(12, []int{2, 100, 120, 600, 12, 12}, []int{120, 600, 12, 12})
}
