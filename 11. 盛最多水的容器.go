
// 解题思路
// 有点像贪心算法的感觉，但是应该不是，这几个概念我目前还没有弄清楚
// 使用双指针进行前后移动，计算两个指针之间的面积
// 当一个指针比另一个指针对应的值要小时，指针移动，然后重新计算
// 这里有点贪心的意思，就是想尝试最大值去进行计算
func maxArea(height []int) int {
	i := 0
	j := len(height) - 1
	var maxValue int
	for i < j {
		area := (j - i) * minHeight(height[i], height[j])
		if area > maxValue {
			maxValue = area
		}
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return maxValue
}

func minHeight(a, b int) int {
	if a < b {
		return a
	}
	return b
}
