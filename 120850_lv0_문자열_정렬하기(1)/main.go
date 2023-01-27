package main

import (
	"fmt"
	"sort"
	"time"
)

func solution(my_string string) []int {

	answer := []int{}
	for _, r := range my_string {
		if '0' <= r && r <= '9' {
			answer = append(answer, int(r-'0'))
		}
	}

	sort.Ints(answer)

	return answer
}

func report(my_string string, desireResult []int) {

	input := fmt.Sprintf("Input:[%v]", my_string)

	start := time.Now()

	result := solution(my_string)

	duration := time.Since(start)

	isSame := true
	if len(result) != len(desireResult) {
		isSame = false
	} else {
		for i, _ := range result {
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
	report("hi12392", []int{1, 2, 2, 3, 9})
	report("p2o4i8gj2", []int{2, 2, 4, 8})
	report("abcde0", []int{0})
}
