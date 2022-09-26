package test1

func IsPrime(num int) bool {
	if num == 2 {
		return true
	}
	for i := 2; i*i <= num; i++ {
		if num % i == 0 {
			return false
		}
	}
	return true
}

func Print() []int {
	ans := make([]int, 0)
	for i := 2; i <= 100; i++ {
		if IsPrime(i) {
			ans = append(ans, i)
		}
	}
	return ans
}