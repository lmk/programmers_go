package main

import (
	"fmt"
	"time"
)

func solution(s string, n int) string {
	l := "abcdefghijklmnopqrstuvwxyz"
	u := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	result := ""
	for _, c := range s {
		if 'a' <= c && c <= 'z' {
			i := (int(c-'a') + n) % len(l)
			c = rune(l[i])
		} else if 'A' <= c && c <= 'Z' {
			i := (int(c-'A') + n) % len(u)
			c = rune(u[i])
		}

		result += string(c)
	}

	return result
}
func report(s string, n int, desireResult string) {

	start := time.Now()

	result := solution(s, n)

	duration := time.Since(start)

	if result == desireResult {
		fmt.Printf("[O] ")
	} else {
		fmt.Printf("[X] ")
	}

	fmt.Printf("Input:[%v], result:[%v], duration:%v\n", s, result, duration)
}

func main() {

	report("AB", 1, "BC")
	report("z", 1, "a")
	report("a B z", 4, "e F d")
}
