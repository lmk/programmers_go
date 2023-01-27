package main

import (
	"fmt"
	"sort"
	"time"
)

func far(a, b int) int {
	if a > b {
		return a - b
	}

	return b - a
}

func solution(numlist []int, n int) []int {

	sort.Slice(numlist, func(i, j int) bool {

		di := far(n, numlist[i])
		dj := far(n, numlist[j])

		if di < dj {
			return true
		} else if di > dj {
			return false
		}

		return numlist[i] > numlist[j]
	})

	return numlist
}

func report(numlist []int, n int, desireResult []int) {

	input := fmt.Sprintf("Input:[%v, %v]", numlist, n)

	start := time.Now()

	result := solution(numlist, n)

	duration := time.Since(start)

	isSame := true
	if result[0] != desireResult[0] || result[1] != desireResult[1] {
		isSame = false
	}

	if isSame {
		fmt.Printf("[O] %s result:[%v] duration:%v\n", input, result, duration)
	} else {
		fmt.Printf("[X] %s result:[%v]->[%v] duration:%v\n", input, result, desireResult, duration)
	}
}

func main() {
	report([]int{1, 2, 3, 4, 5, 6}, 4, []int{4, 5, 3, 6, 2, 1})
	report([]int{10000, 20, 36, 47, 40, 6, 10, 7000}, 30, []int{36, 40, 20, 47, 10, 6, 7000, 10000})
}
