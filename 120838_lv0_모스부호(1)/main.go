package main

import (
	"fmt"
	"strings"
	"time"
)

func solution(letter string) string {
	morses := map[string]string{
		".-": "a", "-...": "b", "-.-.": "c", "-..": "d", ".": "e", "..-.": "f", "--.": "g", "....": "h", "..": "i", ".---": "j", "-.-": "k", ".-..": "l", "--": "m", "-.": "n", "---": "o", ".--.": "p", "--.-": "q", ".-.": "r", "...": "s", "-": "t", "..-": "u", "...-": "v", ".--": "w", "-..-": "x", "-.--": "y", "--..": "z"}

	answer := ""
	for _, m := range strings.Split(letter, " ") {
		answer += morses[m]
	}

	return answer
}

func report(letter string, desireResult string) {

	input := fmt.Sprintf("Input:[%v]", letter)

	start := time.Now()

	result := solution(letter)

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
	report(".... . .-.. .-.. ---", "hello")
	report(".--. -.-- - .... --- -.", "python")
}
