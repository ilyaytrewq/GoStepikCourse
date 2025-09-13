package wordcnt

func WordCount(text string) int {
	cnt := 0
	for _, c := range text {
		if c == ' ' {
			cnt++
		}
	}
	return cnt + 1
}
