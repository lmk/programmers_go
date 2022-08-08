package main

import (
	"fmt"
	"time"
)

func solution(arr []int) float64 {

	sum := 0

	for _, n := range arr {
		sum += n
	}

	return float64(sum) / float64(len(arr))
}

func report(arr []int, desireResult float64) {

	start := time.Now()

	result := solution(arr)

	duration := time.Since(start)

	if result == desireResult {
		fmt.Printf("[O] ")
	} else {
		fmt.Printf("[X] ")
	}

	fmt.Printf("Input:[%v], result:[%v], duration:%v\n", arr, result, duration)
}

func main() {

	report([]int{1, 2, 3, 4}, 2.5)
	report([]int{5, 5}, 5)
}
