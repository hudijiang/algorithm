package easy

import "fmt"

//给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
//
// 你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。
//
// 示例:
//
// 给定 nums = [2, 7, 11, 15], target = 9
//
//因为 nums[0] + nums[1] = 2 + 7 = 9
//所以返回 [0, 1]
//
// Related Topics 数组 哈希表
//leetcode submit region begin(Prohibit modification and deletion)

func twoSum(nums []int, target int) []int {
	if len(nums) < 2 {
		return []int{0, 0}
	}
	maps := map[int]int{}
	fmt.Println(maps[2])
	for i := 0; i < len(nums); i++ {
		_, exists := maps[target-nums[i]]
		if exists {
			return []int{i, maps[target-nums[i]]}
		} else {
			maps[nums[i]] = i
		}
	}
	return []int{0, 0}
}

//leetcode submit region end(Prohibit modification and deletion)