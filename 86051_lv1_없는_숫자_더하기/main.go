package main

import (
	"fmt"
	"time"
)

func solution(numbers []int) int {
	nums := map[int]int{
		1: 1, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8, 9: 9,
	}

	for _, n := range numbers {
		nums[n] = 0
	}

	sum := 0
	for _, v := range nums {
		sum += v
	}

	return sum
}

func report(numbers []int, desireResult int) {

	start := time.Now()

	result := solution(numbers)

	duration := time.Since(start)

	if result == desireResult {
		fmt.Printf("[O] ")
	} else {
		fmt.Printf("[X] ")
	}

	fmt.Printf("Input:[%v], result:[%v], duration:%v\n", numbers, result, duration)
}

func main() {
	report([]int{1, 2, 3, 4, 6, 7, 8, 0}, 14)
	report([]int{5, 8, 4, 0, 6, 7, 9}, 6)
}
