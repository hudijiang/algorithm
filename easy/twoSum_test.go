package easy

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	result := twoSum(nums, 9)
	fmt.Println(result)
}
