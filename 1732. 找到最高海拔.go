
func largestAltitude(gain []int) int {
	// lastest 记录数组从一开始到当前位置的和
	var lastest int
	// maxHeight 记录 lastest 的最大值，lastest 默认为 0
	var maxHeight int
	lenNums := len(gain)
	for i := 0; i < lenNums; i++ {
		lastest += gain[i]
		if lastest > maxHeight {
			maxHeight = lastest
		}
	}
	return maxHeight
}
