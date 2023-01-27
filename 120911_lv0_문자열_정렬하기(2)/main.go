package main

import (
	"fmt"
	"sort"
	"time"
)

func solution(my_string string) string {

	offset := 'a' - 'A'
	arr := []rune{}

	for _, r := range my_string {
		if 'A' <= r && r <= 'Z' {
			r = r + offset
		}

		arr = append(arr, r)
	}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	return string(arr)
}

func report(my_string string, desireResult string) {

	input := fmt.Sprintf("Input:[%v]", my_string)

	start := time.Now()

	result := solution(my_string)

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
	report("Bcad", "abcd")
	report("heLLo", "ehllo")
	report("Python", "hnopty")
}
