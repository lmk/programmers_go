package main

import (
	"fmt"
	"strings"
	"time"
)

func solution(n int, m int) {
	line := strings.Repeat("*", n)
	for m > 0 {
		fmt.Println(line)
		m--
	}
}

func report(n int, m int) {

	start := time.Now()

	solution(n, m)

	duration := time.Since(start)

	fmt.Printf("Input:[%v, %v], duration:%v\n", n, m, duration)
}

func main() {

	report(5, 6)
	report(1, 2)
}
