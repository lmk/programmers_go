package main

import (
	"fmt"
	"time"
)

func solution(dot []int) int {
	if dot[0] > 0 && dot[1] > 0 {
		return 1
	} else if dot[0] < 0 && dot[1] > 0 {
		return 2
	} else if dot[0] < 0 && dot[1] < 0 {
		return 3
	} else if dot[0] > 0 && dot[1] < 0 {
		return 4
	}
	return 0
}

func report(num_list []int, desireResult int) {

	input := fmt.Sprintf("Input:[%v]", num_list)

	start := time.Now()

	result := solution(num_list)

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
	report([]int{2, 4}, 1)
	report([]int{-7, 9}, 2)
}
