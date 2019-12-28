package easy

func reverse(x int) int {
	//结果值
	result := 0

	for x != 0 {
		pop := x % 10
		x = x / 10
		result = result*10 + pop
		//如果溢出则返回
		if !(-(1<<31) <= result && result <= (1<<31)-1) {
			return 0
		}
	}
	return result
}
