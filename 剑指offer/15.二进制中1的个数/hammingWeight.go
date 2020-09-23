package hammingWeight

import (
	"fmt"
	"strings"
)

/*
请实现一个函数，输入一个整数，输出该数二进制表示中 1 的个数。例如，把 9 表示成二进制是 1001，有 2 位是 1。因此，如果输入 9，则该函数输出 2。
*/

func HammingWeight(num uint32) int {
	cnt := 0
	for ; num > 0; num >>= 1 {
		if num&1 == 1 {
			cnt++
		}
	}
	return cnt
}

func HammingWeight1(num uint32) int {
	return strings.Count(fmt.Sprintf("%b", num), "1")
}
