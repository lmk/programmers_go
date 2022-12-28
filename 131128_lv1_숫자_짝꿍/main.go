package main

import (
	"fmt"
	"sort"
	"time"
)

func onlyZero(s []rune) bool {

	for _, r := range s {
		if r != '0' {
			return false
		}
	}

	return true
}

func sortToStr(s []rune) string {

	sort.Slice(s, func(i int, j int) bool {
		return s[i] > s[j]
	})

	return string(s)
}

func solution(X string, Y string) string {

	p := []rune{}

	xc := make([]int, 10)
	yc := make([]int, 10)

	for _, r := range X {
		xc[r-'0']++
	}

	for _, r := range Y {
		yc[r-'0']++
	}

	for i := range xc {

		if yc[i] == 0 || xc[i] == 0 {
			continue
		}

		count := 0
		if xc[i] > yc[i] {
			count = yc[i]
		} else {
			count = xc[i]
		}

		for j := 0; j < count; j++ {
			p = append(p, rune('0'+i))
		}
	}

	if len(p) == 0 {
		return "-1"
	}

	if onlyZero(p) {
		return "0"
	}

	return sortToStr(p)
}

func report(X string, Y string, desireResult string) {

	input := fmt.Sprintf("Input:[%v, %v]", X, Y)

	start := time.Now()

	result := solution(X, Y)

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
	report("100", "2345", "-1")
	report("100", "203045", "0")
	report("100", "123450", "10")
	report("12321", "42531", "321")
	report("5525", "1255", "552")
}
