package iteration

func Repeat(char string, num int) string {
	res := ""
	for i := 0; i < num; i++ {
		res += char
	}
	return res
}
