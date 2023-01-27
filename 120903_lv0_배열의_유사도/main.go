package main

import (
	"fmt"
	"sort"
	"time"
)

func solution(numbers []int) int {

	sort.Slice(numbers, func(i int, j int) bool {
		return numbers[i] > numbers[j]
	})

	return numbers[0] * numbers[1]
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
	report([]int{1, 2, 3, 4, 5}, 20)
	report([]int{0, 31, 24, 10, 1, 9}, 744)
}
