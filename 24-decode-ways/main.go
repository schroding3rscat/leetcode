package main

import "fmt"

func main() {
	s := "12"
	fmt.Println(numDecodings(s))

	s = "226"
	fmt.Println(numDecodings(s))

	s = "2262"
	fmt.Println(numDecodings(s))

	s = "1201234"
	fmt.Println(numDecodings(s))
}

func numDecodings(s string) int {
	l := len(s)
	if l == 0 {
		return 0
	}
	if l == 1 {
		if s[0] == '0' {
			return 0
		}
		return 1
	}
	if l >= 2 && s[0] == '0' {
		return 0
	}

	res := make([]int, l)
	res[0] = 1

	for i := 1; i < l; i++ {
		if s[i] != '0' {
			res[i] = res[i-1]
		}
		if s[i-1] == '0' {
			continue
		}
		if isOK(s[i-1], s[i]) {
			if i == 1 {
				res[i] += res[i-1]
			} else {
				res[i] += res[i-2]
			}
		}
	}

	// "2" => 1
	// "22" => "2"+"22" => 2
	// "226" => "2"+"22"+"26"=>3
	// "2262" => "2"+"22"+"26"+"4"=>4   2 2 6 4, 2 26 4, 22 6 4,

	// "1201234" => 1 20 1 2 3 4, 1 20 12 3 4, 1 20 1 23 4

	return res[l-1]
}

func isOK(s1, s2 byte) bool {
	if s1 == 0 {
		return false
	}
	if s1 == '1' || s1 == '2' && s2 <= '6' {
		return true
	}
	return false
}
