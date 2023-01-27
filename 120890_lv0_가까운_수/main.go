package main

import (
	"fmt"
	"math"
	"time"
)

func solution(array []int, n int) int {
	answer := array[0]
	for _, e := range array {
		if math.Abs(float64(answer-n)) > math.Abs(float64(e-n)) {
			answer = e
		} else if math.Abs(float64(answer-n)) == math.Abs(float64(e-n)) {
			if answer > e {
				answer = e
			}
		}
	}

	return answer
}

func report(array []int, num int, desireResult int) {

	input := fmt.Sprintf("Input:[%v, %v]", array, num)

	start := time.Now()

	result := solution(array, num)

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
	report([]int{3, 10, 28}, 20, 28)
	report([]int{10, 11, 12}, 13, 12)
}
