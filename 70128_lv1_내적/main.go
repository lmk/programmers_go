package main

import (
	"fmt"
	"time"
)

func solution(a []int, b []int) int {

	sum := 0
	for i, _ := range a {
		sum += a[i] * b[i]
	}

	return sum
}

func report(a []int, b []int, desireResult int) {

	start := time.Now()

	result := solution(a, b)

	duration := time.Since(start)

	if result == desireResult {
		fmt.Printf("[O] ")
	} else {
		fmt.Printf("[X] ")
	}

	fmt.Printf("Input:[%v, %v], result:[%v], duration:%v\n", a, b, result, duration)
}

func main() {
	report([]int{1, 2, 3, 4}, []int{-3, -1, 0, 2}, 3)
	report([]int{-1, 0, 1}, []int{1, 0, -1}, -2)
}
