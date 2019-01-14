package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	palindromes := []int{}

	for i := 999; i > 100; i-- {
		for j := 999; j > 100; j-- {
			mul := i * j
			ok, val := isPalindrome(mul)
			if ok {
				palindromes = append(palindromes, val)
			}
		}
	}
	sort.Ints(palindromes)
	fmt.Println(palindromes[len(palindromes)-1])
}

func isPalindrome(val int) (bool, int) {
	s := strconv.Itoa(val)
	strs := strings.Split(s, "")

	for range strs {
		if strs[0] == strs[len(strs)-1] {
			strs = strs[0 : len(strs)-1]
			strs = strs[1:]
		}
		if len(strs) == 0 || len(strs) == 1 {
			return true, val
		}
	}
	return false, 0
}
