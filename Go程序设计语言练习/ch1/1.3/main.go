package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func Echo() {
	s, seq := "", ""
	for _, arg := range os.Args[1:] {
		s += seq + arg
		seq = " "
	}
	fmt.Println(s)
}

func Echo1() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

func main() {
	t := time.Now()
	Echo()
	fmt.Println(time.Now().Sub(t).String())
	Echo1()
	fmt.Println(time.Now().Sub(t).String())
}
