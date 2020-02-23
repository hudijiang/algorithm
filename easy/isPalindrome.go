package easy

//判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
//
// 示例 1:
//
// 输入: 121
//输出: true
//
//
// 示例 2:
//
// 输入: -121
//输出: false
//解释: 从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。
//
//
// 示例 3:
//
// 输入: 10
//输出: false
//解释: 从右向左读, 为 01 。因此它不是一个回文数。
//
//
// 进阶:
//
// 你能不将整数转为字符串来解决这个问题吗？
// Related Topics 数学

//leetcode submit region begin(Prohibit modification and deletion)
func isPalindrome(x int) bool {
	//如果是负数、或者是个位数、或者个位数是10 则不是回文数
	if x < 0 {
		return false
	}
	original := x
	//翻转后的数值
	result := 0

	for x != 0 {
		n := x % 10
		x = x / 10
		result = result*10 + n
	}
	//如果
	return result == original || original == result/10
}

//leetcode submit region end(Prohibit modification and deletion)
