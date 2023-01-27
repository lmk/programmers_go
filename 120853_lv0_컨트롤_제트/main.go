package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func solution(s string) int {

	sum := 0
	arr := strings.Split(s, " ")
	for i, a := range arr {
		if a == "Z" {
			n, _ := strconv.Atoi(arr[i-1])
			sum -= (n)
		} else {
			n, _ := strconv.Atoi(a)
			sum += n
		}
	}

	return sum
}

func report(my_string string, desireResult int) {

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
	report("1 2 Z 3", 4)
	report("10 20 30 40", 100)
	report("10 Z 20 Z 1", 1)
	report("10 Z 20 Z", 0)
	report("-1 -2 -3 Z", -3)
}
