package main

import (
	"fmt"
	"strconv"
)

func FillMap(m map[string]string) {
	for i := 1; i < 5; i++ {
		s := strconv.Itoa(i)
		m[s] = s
	}
}

func main() {
	m := make(map[string]string)
	m["0"] = "a"
	fmt.Println(m)
	FillMap(m)
	fmt.Println(m)
}
