
// 解题思路，利用map key的特性，把数组转换成map进行统计
// 最终使用内存较多，还需要进行后续优化
func uniqueOccurrences(arr []int) bool {
	maps := make(map[int]int)
	for _, v := range arr {
		maps[v]++
	}
	mapBool := make(map[int]bool)
	for _, v := range maps {
		if mapBool[v] == true {
			return false
		}
		mapBool[v] = true
	}
	return true
}
