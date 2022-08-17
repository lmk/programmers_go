package main

import (
	"fmt"
	"sort"
	"time"
)

/*
var gIndex int

type Bigger []string

func (b Bigger) Len() int {
	return len(b)
}

func (b Bigger) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b Bigger) Less(i, j int) bool {

	if b[i][gIndex] == b[j][gIndex] {
		return b[i] < b[j]
	}

	return b[i][gIndex] < b[j][gIndex]
}

func solution(strings []string, n int) []string {
	gIndex = n

	sort.Sort(Bigger(strings))

	return strings
}
*/

func solution(strings []string, n int) []string {

	sort.Slice(strings, func(i, j int) bool {
		if strings[i][n] == strings[j][n] {
			return strings[i] < strings[j]
		}

		return strings[i][n] < strings[j][n]
	})

	return strings
}

func report(strings []string, n int, desireResult []string) {

	input := fmt.Sprintf("Input:[%v, %v]", strings, n)

	start := time.Now()

	result := solution(strings, n)

	duration := time.Since(start)

	isSame := true
	if len(desireResult) != len(result) {
		isSame = false
	} else {
		for i, _ := range desireResult {
			if desireResult[i] != result[i] {
				isSame = false
				break
			}

		}
	}
	ox := ""
	if isSame == true {
		ox = fmt.Sprintf("[O]")
	} else {
		ox = fmt.Sprintf("[X]")
	}

	fmt.Printf("%s %s result:[%v] duration:%v\n", ox, input, result, duration)
}

func main() {
	report([]string{"sun", "bed", "car"}, 1, []string{"car", "bed", "sun"})
	report([]string{"abce", "abcd", "cdx"}, 2, []string{"abcd", "abce", "cdx"})
}
