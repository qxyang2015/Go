package main

import "fmt"

func main() {
	const N = 1024
	var a [N]int
	x := append(a[:N-1:N], 0, 9)
	fmt.Println(cap(a[:N-1:N]), len(a[:N-1:N]))
	y := append(a[:N:N], 9)
	fmt.Println(cap(a[:N:N]), len(a[:N:N]))
	println(cap(x), cap(y))
}
