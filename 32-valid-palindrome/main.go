package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func main() {
	s := "A man, a plan, a canal: Panama"
	fmt.Println(strconv.FormatBool(isPalindrome(s)))

	s = "0P"
	fmt.Println(strconv.FormatBool(isPalindrome(s)))
}

func isPalindrome(s string) bool {
	if s == "" {
		return true
	}

	l := len(s) - 1
	left := 0
	right := l
	for left < right {
		rl := rune(s[left])
		rr := rune(s[right])

		if !unicode.IsDigit(rl) && !unicode.IsLetter(rl) {
			left++
			continue
		}

		if !unicode.IsDigit(rr) && !unicode.IsLetter(rr) {
			right--
			continue
		}

		if rl >= 48 && rl <= 57 && rl != rr {
			return false
		} else if (rl >= 65 && rl <= 90 || rl >= 97 && rl <= 122) && rl != rr && rl != rr+32 && rl+32 != rr {
			return false
		}

		left++
		right--
	}
	return true
}
