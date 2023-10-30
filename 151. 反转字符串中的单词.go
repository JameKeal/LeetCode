
// 代码首次编写，没有考虑内存消耗问题
func reverseWords(s string) string {
	strList := strings.Split(s, " ")
	newStrList := make([]string, 0)
	for i := range strList {
		if strList[i] == "" {
			continue
		}
		newStrList = append(newStrList, strList[i])
	}
	lenNum := len(newStrList)
	for i := 0; i < lenNum/2; i++ {
		newStrList[i], newStrList[lenNum-i-1] = newStrList[lenNum-i-1], newStrList[i]
	}
	return strings.Join(newStrList, " ")
}
