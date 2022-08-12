package main

import (
	"fmt"
	"strings"
	"time"
)

func solution(s string) string {

	s = strings.ToLower(s)
	result := ""

	i := 0
	for _, r := range s {
		if r == ' ' {
			i = 0
			result += string(r)
			continue
		}

		if i%2 == 0 {
			result += strings.ToUpper(string(r))
		} else {
			result += string(r)
		}

		i++
	}

	return result
}

func report(n string, desireResult string) {

	input := fmt.Sprintf("Input:[%v]", n)

	start := time.Now()

	result := solution(n)

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

	report("try hello world", "TrY HeLlO WoRlD")

}
