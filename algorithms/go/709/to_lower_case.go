package main

func toLowerCase(str string) string {
	runs := make([]rune, len(str))
	for i, c := range str {
		if c >= 65 && c <= 91 {
			runs[i] = c + 32
		} else {
			runs[i] = c
		}
	}
	return string(runs)
}
