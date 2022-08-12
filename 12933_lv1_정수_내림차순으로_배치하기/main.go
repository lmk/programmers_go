package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)

func solution(n int64) int64 {
	s := fmt.Sprintf("%v", n)
	as := strings.Split(s, "")
	sort.Sort(sort.Reverse(sort.StringSlice(as)))
	s = strings.Join(as, "")
	n, _ = strconv.ParseInt(s, 10, 64)

	return n
}

func report(n int64, desireResult int64) {

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

	report(118372, 873211)

}
