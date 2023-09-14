package strings

func Join(s ...string) string {
	str := ""

	for _, val := range s {
		str += val
	}

	return str
}
