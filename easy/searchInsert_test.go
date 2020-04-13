package easy

import (
	"fmt"
	"testing"
)

func TestSearchInsert(t *testing.T) {
	nums := []int{1, 3, 5, 6}
	count := searchInsert(nums[:], 5)
	fmt.Printf("%v,%d\n", nums, count)

	nums = []int{1, 3, 5, 6}
	count = searchInsert(nums[:], 2)
	fmt.Printf("%v,%d\n", nums, count)

	nums = []int{1, 3, 5, 6}
	count = searchInsert(nums[:], 7)
	fmt.Printf("%v,%d\n", nums, count)

	nums = []int{1, 3, 5, 6}
	count = searchInsert(nums[:], 5)
	fmt.Printf("%v,%d\n", nums, count)
}
