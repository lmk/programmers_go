package main

import (
	"fmt"
	"time"
)

func solution(seoul []string) string {
	i := 0
	for i, _ = range seoul {
		if seoul[i] == "Kim" {
			break
		}
	}

	return fmt.Sprintf("김서방은 %d에 있다", i)
}

func report(seoul []string, desireResult string) {

	input := fmt.Sprintf("Input:[%v]", seoul)

	start := time.Now()

	result := solution(seoul)

	duration := time.Since(start)

	ox := ""
	if result == desireResult {
		ox = fmt.Sprintf("[O]")
	} else {
		ox = fmt.Sprintf("[X]")
	}

	fmt.Printf("%s %s result:[%v] duration:%v\n", ox, input, result, duration)
}

func main() {
	report([]string{"Jane", "Kim"}, "김서방은 1에 있다")
}
