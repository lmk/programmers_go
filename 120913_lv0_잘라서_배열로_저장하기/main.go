package main

import (
	"fmt"
	"time"
)

func solution(my_str string, n int) []string {

	arr := []string{}
	for i := 0; i < len(my_str); i += n {
		max := i + n
		if max > len(my_str) {
			max = len(my_str)
		}
		arr = append(arr, my_str[i:max])
	}

	return arr
}

func report(my_str string, n int, desireResult []string) {

	input := fmt.Sprintf("Input:[%v, %v]", my_str, n)

	start := time.Now()

	result := solution(my_str, n)

	duration := time.Since(start)

	isSame := true
	if len(result) != len(desireResult) {
		isSame = false
	} else {

		for i := range result {
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
	report("abc1Addfggg4556b", 6, []string{"abc1Ad", "dfggg4", "556b"})
	report("abcdef123", 3, []string{"abc", "def", "123"})
}
