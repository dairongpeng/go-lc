package main

// minLight 给定一个由'X'和'.'组成的居民楼路径。要照亮所有居民楼，返回最少需要几盏灯
func minLight(road string) int {
	str := []byte(road)
	// index从0出发
	index := 0
	// 当前灯的个数
	light := 0
	for index < len(str) {
		// 当前i位置是X，直接跳到下一个位置做决定
		if str[index] == 'X' {
			index++
		} else { // i 位置是 . 不管i+1是X还是.当前区域需要放灯
			light++
			// 接下来没字符了，遍历结束
			if index+1 == len(str) {
				break
			} else {
				// 如果i+1位置是X，在i位置放灯，去i+2位置做决定
				if str[index+1] == 'X' {
					index = index + 2
				} else { // i位置是. i+1也是. 那么不管i+2是什么，都在i+1位置放灯，到i+3去做决定
					index = index + 3
				}
			}
		}
	}
	return light
}
