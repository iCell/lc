// 本题中，可以认为 dp[i][j] 代表从 i-j 个石头堆中，Alex 比 Lee 多的石头数，若是正的则 Alex 赢
// 先考虑最后一步，最后一步肯定只有一个石头堆，那么谁先手谁就赢，dp[i][i] = piles[i]
// 由于每次取石头只能从最左边或者最右边取，因此要不取 i 堆要不取 j 堆，也就是说 i-j 是好多堆石头的情况同最后只有 2 堆石头的情况是一致的，这里指考虑剩下 2 堆石头，由于最后剩下 2 堆，可以是 Alex 先手，也可以是 Lee 先手，每次取的时候都是最优解，所以 piles[0]-dp[1][1] 代表 Alex 先手取走第一批石头得分 piles[0]，而 Lee 之后取走剩下的一堆石头得分 dp[1][1]，反过来 piles[1]-dp[0][0] 代表 Lee 先手而 Alex 后手
// 因此若是多堆石头的情况，若是先手取走第 i 堆，后手取走剩下的 [i+1][j] 堆，则 piles[i]-dp[i+1][j]，也可以是先取走第 j 堆，然后紧接着下一个人取走第 [i][j-1] 堆
// 所以 dp[i][j] = max(piles[i] - dp[i+1][j], piles[j] - dp[i][j-1])

func stoneGame(piles []int) bool {
	dp := make([][]int, len(piles), len(piles))
	for i := range dp {
		dp[i] = make([]int, len(piles), len(piles))
		dp[i][i] = piles[i]
	}

	for j := 1; j < len(dp); j++ {
		for i := 0; i < len(dp)-j; i++ {
			dp[i][j] = max(piles[i]-dp[i+1][j], piles[j]-dp[i][j-1])
		}
	}

	return dp[0][len(piles)-1] > 0
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
