package main

import (
	"fmt"
	"time"
)

func solution(my_string string, num1 int, num2 int) string {
	if num1 > num2 {
		num1, num2 = num2, num1
	}

	return my_string[:num1] + my_string[num2:num2+1] + my_string[num1+1:num2] + my_string[num1:num1+1] + my_string[num2+1:]
}

func report(my_string string, num1 int, num2 int, desireResult string) {

	input := fmt.Sprintf("Input:[%v, %v, %v]", my_string, num1, num2)

	start := time.Now()

	result := solution(my_string, num1, num2)

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
	report("hello", 1, 2, "hlelo")
	report("I love you", 3, 6, "I l veoyou")
}
