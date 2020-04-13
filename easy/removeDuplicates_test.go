package easy

import (
	"fmt"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	nums := []int{1, 1, 2}
	fmt.Printf("%T", nums)
	removeDuplicates(nums[:])
}
