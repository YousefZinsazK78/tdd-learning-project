package iterate

func Repeat(chars string, num int) string {
	var char string
	for i := 0; i < num; i++ {
		char += chars
	}
	return char
}
