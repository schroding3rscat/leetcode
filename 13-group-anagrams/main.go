package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	res := groupAnagrams(strs)
	fmt.Printf("%+v\n", res)

	strs = []string{"", "", ""}
	res = groupAnagrams(strs)
	fmt.Printf("%+v\n", res)
}

func groupAnagrams(strs []string) [][]string {
	if len(strs) <= 1 {
		return [][]string{strs}
	}

	l := len(strs)
	hash := map[string][]string{}

	for i := 0; i <= l-1; i++ {
		sorted := stringSort(strs[i])
		list, ok := hash[sorted]
		if ok {
			hash[sorted] = append(list, strs[i])
		} else {
			hash[sorted] = []string{strs[i]}
		}
	}

	res := [][]string{}
	for _, v := range hash {
		res = append(res, v)
	}

	return res
}

func stringSort(s string) string {
	r := []rune(s)
	sort.Slice(r, func(i, j int) bool { return r[i] < r[j] })
	return string(r)
}
