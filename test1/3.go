package test1

func CountChar(str string) map[byte]int {
	ans := make(map[byte]int, 0)
	for i := 0; i < len(str); i++ {
		ans[str[i]]++
	}
	return ans
}
