
// 解题思路
// 1、左侧之和等于右侧之和，这就意味着，整个数组总和，减去下标数，肯定是个偶数，也就是可以被2除尽
// 换言之，不能被2除尽，就没有成立条件的机会，可以直接pass
// 2、数组总和减去下标数，除以2后的值，换一种说法，其实就是去计算，从数组开头到当前下标的数之和，
// 与除以2后的值是否相等，相等就是我们要的下标位置
// 3、累积计算数组前i个元素之和，可以利用DP的思路，用一个变量去记录，这样我们就只用在上一个和之上，
// 再加上当前元素，就能获得累积计算数组前i个元素之和，不用重复从0开始计算
func pivotIndex(nums []int) int {
	// 首先计算出整个数组的数值总和
	var total int
	for i := 0; i < len(nums); i++ {
		total += nums[i]
	}

	var lastest int
	for i := 0; i < len(nums); i++ {
		if (total-nums[i])%2 == 0 {
			if (total-nums[i])/2 == lastest {
				return i
			}
		}
		lastest += nums[i]
	}
	return -1
}
