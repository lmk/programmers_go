package main

import (
	"fmt"
	"time"
)

func solution(number []int) int {
	answer := 0
	for i := 0; i < len(number)-2; i++ {
		for j := i + 1; j < len(number)-1; j++ {
			for k := j + 1; k < len(number); k++ {
				if number[i]+number[j]+number[k] == 0 {
					answer++
				}
			}
		}
	}
	return answer
}

func report(number []int, desireResult int) {

	input := fmt.Sprintf("Input:[%v]", number)

	start := time.Now()

	result := solution(number)

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
	report([]int{-2, 3, 0, 2, -5}, 2)
	report([]int{-3, -2, -1, 0, 1, 2, 3}, 5)
	report([]int{-1, 1, -1, 1}, 0)
}
