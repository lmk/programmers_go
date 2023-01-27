package main

import (
	"fmt"
	"time"
)

func solution(num_list []int) []int {
	for i := 0; i < len(num_list)/2; i++ {
		num_list[i], num_list[len(num_list)-1-i] = num_list[len(num_list)-1-i], num_list[i]
	}

	return num_list
}

func report(num_list []int, desireResult []int) {

	input := fmt.Sprintf("Input:[%v]", num_list)

	start := time.Now()

	result := solution(num_list)

	duration := time.Since(start)

	isSame := true
	if len(num_list) != len(desireResult) {
		isSame = false
	} else {
		for i, _ := range num_list {
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
	report([]int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1})
	report([]int{1, 1, 1, 1, 1, 2}, []int{2, 1, 1, 1, 1, 1})
	report([]int{1, 0, 1, 1, 1, 3, 5}, []int{5, 3, 1, 1, 1, 0, 1})
}
