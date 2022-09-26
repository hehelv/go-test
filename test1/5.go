package test1

func PrintSlice() []int {
	res := make([]int, 0)
	for i := 1; i <= 50; i++ {
		if i % 3 == 0 {
			continue
		}
		res = append(res, i)
	}
	res = append(res, 666)
	return res
}
