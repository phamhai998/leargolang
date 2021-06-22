package main

import (
	"fmt"
)

/*
Given a string s, find the length of the longest substring without repeating characters.
Input: s = "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.
*/

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	// kiem tra neu no <= 1 thi se in ngay ra ket qua
	if n <= 1 {
		return n
	}
	//
	myMap := make(map[string]int, 0)
	//
	max := 0
	start := 0
	//tim vi tri roi so sanh de lay khoang cac lon nhat
	for k, v := range s {
		c := string(v)
		i, ok := myMap[c]
		// tim kiem va lay vi tri trong map
		if ok && i >= start {
			start = i + 1
		}
		myMap[c] = k
		if k-start+1 > max {
			max = k - start + 1
		}
	}
	return max
}

func main() {
	s := "acdabcd"
	fmt.Print(lengthOfLongestSubstring(s))
	// lengthOfLongestSubstring(s)
}
