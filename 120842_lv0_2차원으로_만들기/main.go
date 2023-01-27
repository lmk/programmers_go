package main

import (
	"fmt"
	"time"
)

func solution(num_list []int, n int) [][]int {

	r := len(num_list) / n
	answer := make([][]int, r)

	for i := 0; i < r; i++ {
		answer[i] = num_list[i*n : i*n+n]
	}

	return answer
}

func report(num_list []int, num int, desireResult [][]int) {

	input := fmt.Sprintf("Input:[%v, %v]", num_list, num)

	start := time.Now()

	result := solution(num_list, num)

	duration := time.Since(start)

	isSame := true
	if len(result) != len(desireResult) {
		isSame = false
	} else {
		for i, sub := range result {
			for j, _ := range sub {
				if result[i][j] != desireResult[i][j] {
					isSame = false
					break
				}
			}
			if isSame == false {
				break
			}
		}
	}

	if isSame {
		fmt.Printf("[O] %s result:[%v] duration:%v\n", input, result, duration)
	} else {
		fmt.Printf("[X] %s result:[%v]->[%v] duration:%v\n", input, result, desireResult, duration)
	}
}

func main() {
	report([]int{1, 2, 3, 4, 5, 6, 7, 8}, 2, [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}})
	report([]int{100, 95, 2, 4, 5, 6, 18, 33, 948}, 3, [][]int{{100, 95, 2}, {4, 5, 6}, {18, 33, 948}})
}
