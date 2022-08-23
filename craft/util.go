package craft

// 判断是否为数字
func isDigital(s byte) bool {
	if s >= '0' && s <= '9' {
		return true
	}
	return false
}

// 判断是否为字母
func isAlpha(s byte) bool {
	if (s >= 'a' && s <= 'z') || (s >= 'A' && s <= 'Z') {
		return true
	}
	return false
}
