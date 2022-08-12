package main

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

func solution(s string) string {

	ss := strings.Split(s, "")
	sort.Sort(sort.Reverse(sort.StringSlice(ss)))

	return strings.Join(ss, "")
}
func report(s string, desireResult string) {

	input := fmt.Sprintf("Input:[%v]", s)

	start := time.Now()

	result := solution(s)

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
	report("Zbcdefg", "gfedcbZ")
}
