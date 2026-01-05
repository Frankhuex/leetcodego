package leetcode

func longestPalindrome(s string) string {
	dp := make([][]bool, len(s))
	ans := s[0:1]
	for i := 0; i < len(s); i++ { //slice length = i+1
		dp[i] = make([]bool, len(s)-i)
		for j := 0; j < len(s)-i; j++ {
			if i == 0 {
				dp[i][j] = true
			} else if i == 1 {
				dp[i][j] = (s[j] == s[j+1])
			} else {
				dp[i][j] = (s[j] == s[j+i] && dp[i-2][j+1])
			}
			if dp[i][j] {
				ans = s[j : j+i+1]
			}
		}
	}
	return ans
}
