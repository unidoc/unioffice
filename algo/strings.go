package algo

func RepeatString(s string, cnt int) string {
	if cnt <= 0 {
		return ""
	}
	buf := make([]byte, len(s)*cnt)
	sb := []byte(s)
	for i := 0; i < cnt; i++ {
		copy(buf[i:], sb)
	}
	return string(buf)
}
