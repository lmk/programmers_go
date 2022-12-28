package main

import (
	"fmt"
	"strings"
	"time"
)

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func solution(food []int) string {

	food = food[1:]

	left := ""
	for i, p := range food {
		left += strings.Repeat(fmt.Sprintf("%d", i+1), p/2)
	}

	return left + "0" + Reverse(left)
}

func report(food []int, desireResult string) {

	input := fmt.Sprintf("Input:[%v]", food)

	start := time.Now()

	result := solution(food)

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
	report([]int{1, 3, 4, 6}, "1223330333221")
	report([]int{1, 7, 1, 2}, "111303111")
}
