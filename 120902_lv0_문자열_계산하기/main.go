package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func atoi(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func solution(my_string string) int {

	num := 0
	op := ""
	a := strings.Split(my_string, " ")

	for i := 0; i < len(a); i++ {
		if a[i] == "+" || a[i] == "-" {
			op = a[i]
		} else {
			n := atoi(a[i])
			if op == "+" {
				num += n
			} else if op == "-" {
				num -= n
			} else {
				num = n
			}
			op = ""
		}
	}

	return num
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
	report("3 + 4", 7)
}
