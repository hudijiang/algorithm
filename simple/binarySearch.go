package simple

import "fmt"

// 二分查找法
func Rank(key int, arr []int) int {
	count := 0
	// 声明数据查找范围 startIndex 开始位置 endIndex 结束位置
	startIndex, endIndex := 0, len(arr)
	for {
		count++
		if startIndex > endIndex {
			break
		}
		// 求中间位置
		middle := startIndex + (endIndex-startIndex)/2

		if key < arr[middle] {
			endIndex = middle - 1
		} else if key > arr[middle] {
			startIndex = middle + 1
		} else {
			fmt.Println("共查找了", count, "次")
			return middle
		}
	}
	fmt.Println("共查找了", count, "次")
	return -1
}

func RecursionRank(key int,arry []int) int{
	return Ranks(key,0,len(arry),arry)
}

//递归式二分查找
func Ranks(key,start,end int,arr []int) int{
	//如果不存在，则直接返回
	if start > end {
		return -1
	}
	//取中间索引位置
	mid := start + (end-start)/2
	//如果需要查找的数据比中间位置小，则继续递归
	if key < arr[mid] {
		return Ranks(key,start,mid-1,arr)
	} else if key > arr[mid] {
		return Ranks(key,mid+1,end,arr)
	}
	return mid
}
