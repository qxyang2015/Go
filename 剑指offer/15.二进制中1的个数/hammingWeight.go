package hammingWeight

import (
	"fmt"
	"strings"
)

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
