package main

import (
	"fmt"
	"sort"
	"time"
)

func solution(array []int) int {

	sort.Slice(array, func(i int, j int) bool {
		return array[i] < array[j]
	})

	return array[len(array)/2]
}

func report(array []int, desireResult int) {

	input := fmt.Sprintf("Input:[%v]", array)

	start := time.Now()

	result := solution(array)

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
	report([]int{1, 2, 7, 10, 11}, 7)
	report([]int{9, -1, 0}, 0)
}
