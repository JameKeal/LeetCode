
// 解题思路，简单来说，就是把数组转换成map，利用map key的特性，方便快速查找
// 这是脑子第一反应想出来的解题思路
func findDifference(nums1 []int, nums2 []int) [][]int {
	map1 := make(map[int]int, len(nums1))
	for i := 0; i < len(nums1); i++ {
		map1[nums1[i]]++
	}
	map2 := make(map[int]int, len(nums2))
	for i := 0; i < len(nums2); i++ {
		map2[nums2[i]]++
		if _, ok := map1[nums2[i]]; ok {
			delete(map1, nums2[i])
		}
	}

	for i := 0; i < len(nums1); i++ {
		if _, ok := map2[nums1[i]]; ok {
			delete(map2, nums1[i])
		}
	}

	array1 := make([]int, 0, len(map1))
	for key := range map1 {
		array1 = append(array1, key)
	}
	array2 := make([]int, 0, len(map2))
	for key := range map2 {
		array2 = append(array2, key)
	}
	result := make([][]int, 2)
	result[0] = array1
	result[1] = array2
	return result
}
