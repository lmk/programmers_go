package main

import (
	"fmt"
	"time"
)

func solution(arr []int) []int {
	result := []int{}

	lo := 0
	for i, n := range arr {
		if arr[lo] > n {
			lo = i
		}
	}

	result = append(arr[:lo], arr[lo+1:]...)

	if len(result) == 0 {
		result = append(result, -1)
	}

	return result
}

func report(arr []int, desireResult []int) {

	input := fmt.Sprintf("Input:[%v]", arr)

	start := time.Now()

	result := solution(arr)

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
	report([]int{4, 3, 2, 1}, []int{4, 3, 2})
	report([]int{10}, []int{-1})
}
