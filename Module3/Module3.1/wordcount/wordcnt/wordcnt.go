package wordcnt

func WordCount(text string) int {
	cnt := 0
	for _, c := range text {
		if c == ' ' {
			cnt++
		}
	}
	if cnt == len(text) {
		return 0
	}
	return cnt + 1
}
