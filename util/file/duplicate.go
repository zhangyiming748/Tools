package file

func Duplicate(src, dst string) {
	var passwd = map[string]bool{}
	for _, v := range ReadLine(src) {
		if _, ok := passwd[v]; ok {
			continue
		} else {
			passwd[v] = true
		}
	}

	after := make([]string, 0)
	for k, _ := range passwd {
		after = append(after, k)
	}

	WriteLines(dst, after)
}
