
// 首次提交代码，使用最基础的，最直接的方式进行处理
func canPlaceFlowers(flowerbed []int, n int) bool {
	var count, i, num int
	for i, num = range flowerbed {
		if num == 0 {
			count++
		} else {
			// 连续0被中断，计算下可以种植几朵花
			if count == i {
				plantNums := count / 2
				n -= plantNums
			} else if count >= 3 {
				plantNums := (count - 1) / 2
				n -= plantNums
			}
			// 计算完毕后清理计数
			count = 0
		}
	}

	if count > 0 && count == len(flowerbed) {
		plantNums := count / 2
		if count % 2 == 1 {
			plantNums += 1
		}
		n -= plantNums
	} else if count > 0 && i == len(flowerbed)-1 {
		plantNums := count / 2
		n -= plantNums
	}

	if n <= 0 {
		return true
	}
	return false
}

// 发现问题
// 场景考虑不周全，像连续开头、结尾都是0的时候，一开始没有做处理；全部都为0的时候，也没有想到
