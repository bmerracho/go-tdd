package iteration

func Repeat(char string, cnt int) (rep string) {
	for i := 0; i < cnt; i++ {
		rep += char
	}

	return rep
}
