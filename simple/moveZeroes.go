package simple

import "fmt"

func MoveZeroes(nums []int)  {
	index := 0
	for _,val:=range nums {
		if val!=0 {
			nums[index]=val
			index++
		}
	}
	for i:=index;i<len(nums);i++ {
		nums[i]=0
	}
	fmt.Println(nums)
}
