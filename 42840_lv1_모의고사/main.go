package main

import (
	"fmt"
	"time"
)

func getScore(answers []int, patt []int) int {
	score := 0

	for i, a := range answers {
		if a == patt[i%len(patt)] {
			score++
		}
	}

	return score
}

func solution(answers []int) []int {

	scores := make([]int, 3)

	scores[0] = getScore(answers, []int{1, 2, 3, 4, 5})
	scores[1] = getScore(answers, []int{2, 1, 2, 3, 2, 4, 2, 5})
	scores[2] = getScore(answers, []int{3, 3, 1, 1, 2, 2, 4, 4, 5, 5})

	max := 0
	for _, n := range scores {
		if max < n {
			max = n
		}
	}

	var result []int
	for i, n := range scores {
		if max == n {
			result = append(result, i+1)
		}
	}

	return result
}
func report(answers []int, desireResult []int) {

	start := time.Now()

	result := solution(answers)

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
		fmt.Printf("[O] ")
	} else {
		fmt.Printf("[X] ")
	}

	fmt.Printf("Input:[%v], result:[%v], duration:%v\n", answers, result, duration)
}

func main() {
	report([]int{1, 2, 3, 4, 5}, []int{1})
	report([]int{1, 3, 2, 4, 2}, []int{1, 2, 3})
}
