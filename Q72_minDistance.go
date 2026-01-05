package leetcode

func minDistance(word1 string, word2 string) int {
	dp := make([][]int, len(word1)+1)
	for i := range dp {
		dp[i] = make([]int, len(word2)+1)
		if i == 0 {
			for j := range dp[i] {
				dp[i][j] = j
			}
		} else {
			dp[i][0] = i
		}
	}
	for i, c1 := range word1 {
		for j, c2 := range word2 {
			if c1 == c2 {
				dp[i+1][j+1] = min3(dp[i][j], dp[i][j+1]+1, dp[i+1][j]+1)
			} else {
				dp[i+1][j+1] = min3(dp[i][j], dp[i][j+1], dp[i+1][j]) + 1
			}
		}
	}
	return dp[len(word1)][len(word2)]
}

func min3(a, b, c int) int {
	return min(a, min(b, c))
}
