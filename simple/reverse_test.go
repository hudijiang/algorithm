package simple

import (
	"fmt"
	"testing"
)

func TestReverseString(t *testing.T) {
	ReverseString("1234553ABCDEF")
}

func TestReverseNumber(t *testing.T) {
	//1534236469
	//2147483647
	fmt.Println(ReverseNumber(1534236469))
	fmt.Println(ReverseNumber(-7680))
	fmt.Println(ReverseNumber(878788787777666777))
}
