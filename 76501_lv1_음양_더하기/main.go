package main

import (
	"fmt"
	"time"
)

func solution(absolutes []int, signs []bool) int {

	sum := 0
	for i, n := range absolutes {
		if signs[i] {
			sum += n
		} else {
			sum -= n
		}
	}

	return sum
}

func report(absolutes []int, signs []bool, desireResult int) {

	start := time.Now()

	result := solution(absolutes, signs)

	duration := time.Since(start)

	if result == desireResult {
		fmt.Printf("[O] ")
	} else {
		fmt.Printf("[X] ")
	}

	fmt.Printf("Input:[%v, %v], result:[%v], duration:%v\n", absolutes, signs, result, duration)
}

func main() {
	report([]int{4, 7, 12}, []bool{true, false, true}, 9)
	report([]int{1, 2, 3}, []bool{false, false, true}, 0)
}
