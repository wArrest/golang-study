package iteration

func Repeat(str string, count int) (res string) {
	for i := 0; i < count; i++ {
		res += str
	}
	return
}
