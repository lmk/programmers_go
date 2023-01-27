package main

import (
	"fmt"
	"time"
)

func solution(money int) []int {
	answer := []int{money / 5500, money % 5500}

	return answer
}

func report(money int, desireResult []int) {

	input := fmt.Sprintf("Input:[%v]", money)

	start := time.Now()

	result := solution(money)

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
	report(5500, []int{1, 0})
	report(15000, []int{2, 4000})
}
