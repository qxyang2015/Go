package main

import (
	"fmt"
	"github.com/qxyang2015/Go/sort"
)

func main() {
	fmt.Println("start")
	arrRaw := []int{1, 3, 2, 5, 7, 5}
	fmt.Println("raw :", arrRaw)
	arrSort := sort.BubbleSort2(arrRaw)
	fmt.Println("sort:", arrSort)
	fmt.Println("done!")
}
