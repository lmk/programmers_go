package main

import (
	"fmt"
	"time"
)

func solution(score [][]int) []int {

	ret := make([]int, len(score))

	for i, s := range score {
		sv := float64(s[0]+s[1]) / 2
		ret[i] = 1

		for j, t := range score {
			if i == j {
				continue
			}

			tv := float64(t[0]+t[1]) / 2
			if sv < tv {
				ret[i]++
			}
		}
	}

	return ret
}

func report(score [][]int, desireResult []int) {

	input := fmt.Sprintf("Input:[%v]", score)

	start := time.Now()

	result := solution(score)

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
	report([][]int{{80, 70}, {90, 50}, {40, 70}, {50, 80}}, []int{1, 2, 4, 3})
	report([][]int{{80, 70}, {70, 80}, {30, 50}, {90, 100}, {100, 90}, {100, 100}, {10, 30}}, []int{4, 4, 6, 2, 2, 1, 7})
}
