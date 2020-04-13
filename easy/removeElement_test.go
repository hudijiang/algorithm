package easy

import (
	"fmt"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	count := removeElement(nums[:], 2)
	fmt.Printf("%v,%d\n", nums, count)
	nums = []int{3, 2, 2, 3}
	count = removeElement(nums[:], 3)
	fmt.Printf("%v,%d\n", nums, count)
}
