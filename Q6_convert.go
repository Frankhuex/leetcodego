package leetcode

func convert(s string, numRows int) string {
	if numRows == 1 || len(s) <= numRows {
		return s
	}
	T := 2*numRows - 2
	bytes := make([]byte, 0, len(s))
	for i := 0; i < len(s); i += T {
		bytes = append(bytes, s[i])
	}
	for row := 1; row < numRows-1; row++ {
		bytes = append(bytes, s[row])
		for i := T; i-row < len(s); i += T {
			bytes = append(bytes, s[i-row])
			if i+row < len(s) {
				bytes = append(bytes, s[i+row])
			}
		}
	}
	for i := numRows - 1; i < len(s); i += T {
		bytes = append(bytes, s[i])
	}
	return string(bytes)
}
