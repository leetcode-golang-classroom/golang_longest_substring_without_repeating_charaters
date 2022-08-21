package sol

func lengthOfLongestSubstring(s string) int {
	sLen := len(s)
	start, maxLen := 0, 0
	visitMap := make(map[byte]int)
	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for end := 0; end < sLen; end++ {
		// after start visit same character
		if pos, ok := visitMap[s[end]]; ok && pos >= start {
			start = pos + 1
		}
		visitMap[s[end]] = end
		maxLen = max(maxLen, end-start+1)
		if start+maxLen >= sLen {
			break
		}
	}
	return maxLen
}
