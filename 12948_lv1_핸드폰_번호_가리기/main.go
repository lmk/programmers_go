package main

import (
	"fmt"
	"time"
)

func solution(phone_number string) string {
	length := len(phone_number)
	result := ""
	for i := 0; i < length-4; i++ {
		result += "*"
	}
	return result + phone_number[length-4:length]
}

func report(phone_number string, desireResult string) {

	start := time.Now()

	result := solution(phone_number)

	duration := time.Since(start)

	if result == desireResult {
		fmt.Printf("[O] ")
	} else {
		fmt.Printf("[X] ")
	}

	fmt.Printf("Input:[%v], result:[%v], duration:%v\n", phone_number, result, duration)
}

func main() {

	report("01033334444", "*******4444")
	report("027778888", "*****8888")
}
