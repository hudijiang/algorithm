package easy

import (
	"fmt"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	test := []int{121, -121, 10, 13121, 12321, 0}
	for _, n := range test {
		fmt.Printf("%d is palindrome ? %v\n", n, isPalindrome(n))
	}
}
