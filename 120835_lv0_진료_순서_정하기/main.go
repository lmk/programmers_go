package main

import (
	"fmt"
	"sort"
	"time"
)

type Order struct {
	em int
	i  int
}

func solution(emergency []int) []int {

	result := []Order{}
	for i, e := range emergency {
		result = append(result, Order{e, i})
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].em > result[j].em
	})

	for i, e := range result {
		emergency[e.i] = i + 1
	}

	return emergency
}

func report(emergency []int, desireResult []int) {

	input := fmt.Sprintf("Input:[%v]", emergency)

	start := time.Now()

	result := solution(emergency)

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
	report([]int{3, 76, 24}, []int{3, 1, 2})
	report([]int{1, 2, 3, 4, 5, 6, 7}, []int{7, 6, 5, 4, 3, 2, 1})
	report([]int{30, 10, 23, 6, 100}, []int{2, 4, 3, 5, 1})
}
