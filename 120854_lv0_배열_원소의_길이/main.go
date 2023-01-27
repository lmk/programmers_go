package main

import (
	"fmt"
	"time"
)

func solution(strlist []string) []int {

	answer := make([]int, len(strlist))

	for i, s := range strlist {
		answer[i] = len(s)
	}

	return answer
}

func report(strlist []string, desireResult []int) {

	input := fmt.Sprintf("Input:[%v]", strlist)

	start := time.Now()

	result := solution(strlist)

	duration := time.Since(start)

	isSame := true
	if len(result) != len(desireResult) {
		isSame = false
	} else {
		for i := range desireResult {
			if result[i] != desireResult[i] {
				isSame = false
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
	report([]string{"We", "are", "the", "world!"}, []int{2, 3, 3, 6})
	report([]string{"I", "Love", "Programmers."}, []int{1, 4, 12})
}
