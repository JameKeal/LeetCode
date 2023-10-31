
// 解题思路
// 先将数组进行升序排序，然后双指针一头一尾进行移动
// 两指针之和如果大于k，往前移动尾指针，小于k，往后移动头指针
func maxOperations(nums []int, k int) int {
	sort.Ints(nums)
	i := 0
	j := len(nums) - 1
	var count int
	for i < j {
		if nums[i]+nums[j] > k {
			j--
		} else if nums[i]+nums[j] < k {
			i++
		} else {
			count++
			i++
			j--
		}
	}
	return count
}
