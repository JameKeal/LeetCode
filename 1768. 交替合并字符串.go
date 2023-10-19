
// 首次提交代码，使用最基础的，最直接的方式进行处理
func mergeAlternately(word1 string, word2 string) string {
	lenWord1 := len(word1)
	lenWord2 := len(word2)
	min := lenWord1
	if min > lenWord2 {
		min = lenWord2
	}


	result := make([]byte, 2*min)
	var i, j int
	for ; i < min; i++ {
		result[j] = word1[i]
		result[j+1] = word2[i]
		j += 2
	}

	if lenWord1 != lenWord2 {
		if min == lenWord1 {
			result = append(result, []byte(word2[i:])...)
		} else {
			result = append(result, []byte(word1[i:])...)
		}
	}

	return string(result)
}

/*
发现的问题：
1、针对字符串的操作，遗忘的太多，如果不是使用IDE进行编写代码，很多地方其实都是有语法错误的。比如string是可以使用下标访问具体某个字符的
2、针对数组，一旦创建的时候就定义好长度，后续就不能使用append的方式进行追加，不然中间分配好空间的位置如果没有填充信息，就会把空信息保留
*/
