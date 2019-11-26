package simple

import (
	"fmt"
	"testing"
)

func TestRank(t *testing.T) {
	num := make([]int, 10000000909)
	for i := 0; i < 10000000909; i++ {
		num[i] = i
	}
	var j = Rank(39908702, num)
	fmt.Println(j)
}

func TestRecursionRank(t *testing.T) {
	num := make([]int, 10000000909)
	for i := 0; i < 10000000909; i++ {
		num[i] = i
	}
	var j = RecursionRank(39908702, num)
	fmt.Println(j)
}
