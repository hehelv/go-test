package test1

// IsLeapYear 闰年判断：满足下面其一
// 1. 是4的倍数，不是100的倍数
// 2. 是100的倍数，并且必须是400的倍数
func IsLeapYear(year int) bool {
	if (year % 4 == 0 && year % 100 != 0) || (year % 100 == 0 && year % 400 == 0) {
		return true
	}
	return false
}
