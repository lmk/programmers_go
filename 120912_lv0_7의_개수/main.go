package main

import (
	"fmt"
	"time"
)

func solution(array []int) int {

	c := 0

	for _, a := range array {
		for n := a; n > 0; n /= 10 {
			if n%10 == 7 {
				c++
			}
		}
	}

	return c
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
	report([]int{7, 77, 17}, 4)
	report([]int{10, 29}, 0)
}
