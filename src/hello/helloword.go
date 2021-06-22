package main

import (
	"fmt"
)

func find(num int, nums ...int) {
	fmt.Printf("type of nums is %T\n", nums)
	found := false
	for i, v := range nums {
		if v == num {
			fmt.Println(num, "found at index", i, "in", nums)
			found = true
		}
	}
	if !found {
		fmt.Println(num, "not found in ", nums)
	}
	fmt.Printf("===========\n")
}

func change1(s ...string) {
	s[0] = "GO"
	fmt.Println("==========")
}

func change2(s ...string) {
	s[0] = "GO"
	s = append(s, "playground")
	fmt.Println(s)
	fmt.Println("==========")
}

func main() {
	Welcome := []string{"hello", "world"}
	change2(Welcome...)
	fmt.Println(Welcome)
}
