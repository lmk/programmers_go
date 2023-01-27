package main

import (
	"fmt"
	"sort"
	"time"
)

func solution(n int) []int {

	m := make(map[int]bool)
	for n > 0 {

		if n == 1 {
			break
		}

		i := 2
		for n%i != 0 {
			i++
		}
		m[i] = true
		n = n / i
	}

	arr := []int{}
	for k, _ := range m {
		arr = append(arr, k)
	}

	sort.Ints(arr)

	return arr
}

func report(num int, desireResult []int) {

	input := fmt.Sprintf("Input:[%v]", num)

	start := time.Now()

	result := solution(num)

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
	report(12, []int{2, 3})
	report(17, []int{17})
	report(420, []int{2, 3, 5, 7})
}
