package main

import (
	"fmt"
	"sort"
	"time"
)

func solution(data [][]int, col int, row_begin int, row_end int) int {
	// zero base
	col--
	row_begin--
	row_end--

	// sort
	sort.Slice(data, func(i, j int) bool {
		if data[i][col] == data[j][col] {
			return data[i][0] > data[j][0]
		}

		return data[i][col] < data[j][col]
	})

	sArr := []int{}
	for i := row_begin; i <= row_end; i++ {
		s := 0
		for _, j := range data[i] {
			s += j % (i + 1)
		}
		sArr = append(sArr, s)
	}

	answer := 0
	for i, s := range sArr {
		if i == 0 {
			answer = s
		} else {
			answer = answer ^ s
		}
	}

	return answer
}

func report(data [][]int, col int, row_begin int, row_end int, desireResult int) {

	input := fmt.Sprintf("Input:[%v, %v, %v, %v]", data, col, row_begin, row_end)

	start := time.Now()

	result := solution(data, col, row_begin, row_end)

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
	report([][]int{{2, 2, 6}, {1, 5, 10}, {4, 2, 9}, {3, 8, 3}}, 2, 2, 3, 4)
}
