package simple

import (
	"fmt"
)

func ReverseString(str string) string {
	ch := []rune(str)
	for i:=0;i< len(ch)/2;i++ {
		ch[i],ch[len(ch)-i-1]=ch[len(ch)-i-1],ch[i]
	}
	fmt.Println(string(ch))
	return string(ch)
}



func ReverseNumber(x int) int {
	y := 0
	for x!=0 {
		y = y*10 + x%10
		if !( -(1<<31) <= y && y <= (1<<31)-1) {
			return 0
		}
		x /= 10
	}
	return y
}
