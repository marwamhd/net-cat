package mainhelper

func Itoa(num int) string {
	if num == 0 {
		return "0"
	}
	sign := ""
	if num < 0 {
		sign = "-"
		num = -num
	}
	result := ""
	for num > 0 {
		result = string(rune(num%10+'0')) + result
		num = num / 10
	}
	return sign + result
}
