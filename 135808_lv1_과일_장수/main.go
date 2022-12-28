package main

import (
	"fmt"
	"sort"
	"time"
)

func solution(k int, m int, score []int) int {

	sort.Slice(score, func(i, j int) bool {
		return score[i] > score[j]
	})

	sum := 0
	for i := 0; i < len(score); i += m {

		// 상자에 담기
		if i+m <= len(score) {
			a := score[i : i+m]

			sum += len(a) * a[m-1]
		}
	}

	return sum
}

func report(k int, m int, score []int, desireResult int) {

	input := fmt.Sprintf("Input:[%v, %v, %v]", k, m, score)

	start := time.Now()

	result := solution(k, m, score)

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
	report(3, 4, []int{1, 2, 3, 1, 2, 3, 1}, 8)
	report(4, 3, []int{4, 1, 2, 2, 4, 4, 4, 4, 1, 2, 4, 2}, 33)
}
