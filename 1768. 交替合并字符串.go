
// 首次提交代码，使用最基础的，最直接的方式进行处理
func mergeAlternately(word1 string, word2 string) string {
	lenWord1, lenWord2 := len(word1), len(word2)
	min := lenWord1
	if min > lenWord2 {
		min = lenWord2
	}

	//result := make([]byte, 2*min)
	result := make([]byte, lenWord1+lenWord2)
	var i, j int
	for ; i < min; i, j = i+1, j+2 {
		result[j] = word1[i]
		result[j+1] = word2[i]
	}

	if lenWord1 != lenWord2 {
		if min == lenWord1 {
			//result = append(result, []byte(word2[i:])...)
			for ; i < lenWord2; i, j = i+1, j+1 {
				result[j] = word2[i]
			}
		} else {
			//result = append(result, []byte(word1[i:])...)
			for ; i < lenWord1; i, j = i+1, j+1 {
				result[j] = word1[i]
			}
		}
	}

	return string(result)
}

/*
发现的问题：
1、针对字符串的操作，遗忘的太多，如果不是使用IDE进行编写代码，很多地方其实都是有语法错误的。比如string是可以使用下标访问具体某个字符的
2、针对数组，一旦创建的时候就定义好长度，后续就不能使用append的方式进行追加，不然中间分配好空间的位置如果没有填充信息，就会把空信息保留
3、代码中注释的地方，也是另外一种解题方式，不过由于使用append的方式进行追加，导致内存使用相较数组一开始分配好内存相比，内存使用过多
*/

// 官方解法
func mergeAlternately(word1, word2 string) string {
    n, m := len(word1), len(word2)
    ans := make([]byte, 0, n+m)
    for i := 0; i < n || i < m; i++ {
        if i < n {
            ans = append(ans, word1[i])
        }
        if i < m {
            ans = append(ans, word2[i])
        }
    }
    return string(ans)
}
// 相较之下，解题思路答题相同，但是代码简洁程度，差距不是一星半点，需要自己慢慢磨练下，太久没有写这种基础代码了，但是基础代码反而是最显基础能力的地方，是地基，需要夯实
