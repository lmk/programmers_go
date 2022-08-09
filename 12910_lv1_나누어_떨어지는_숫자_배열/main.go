package main

import (
	"fmt"
	"sort"
	"time"
)

func solution(arr []int, divisor int) []int {
	result := []int{}
	for _, n := range arr {
		if n%divisor == 0 {
			result = append(result, n)
		}
	}

	if len(result) == 0 {
		result = append(result, -1)
	}

	sort.Ints(result)

	return result
}

func report(arr []int, divisor int, desireResult []int) {

	input := fmt.Sprintf("Input:[%v %v]", arr, divisor)

	start := time.Now()

	result := solution(arr, divisor)

	duration := time.Since(start)

	isSame := true
	if len(result) != len(desireResult) {
		isSame = false
	} else {
		for i, v := range result {
			if desireResult[i] != v {
				isSame = false
				break
			}
		}
	}

	ox := ""
	if isSame {
		ox = fmt.Sprintf("[O]")
	} else {
		ox = fmt.Sprintf("[X]")
	}

	fmt.Printf("%s %s result:[%v] duration:%v\n", ox, input, result, duration)
}

func main() {
	report([]int{5, 9, 7, 10}, 5, []int{5, 10})
	report([]int{2, 36, 1, 3}, 1, []int{1, 2, 3, 36})
	report([]int{3, 2, 6}, 10, []int{-1})
}
