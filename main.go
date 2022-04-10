package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

const (
	INF = int(1 << 61)
	MOD = int(1e9 + 7)
	// MOD = 998244353
)

func main() {
	defer _w.Flush()
	var N int
	var S string
	fmt.Fscan(_r, &N, &S)
	ans := Solve(N, S)
	fmt.Fprintf(_w, "%d\n", ans)
}

func Solve(N int, S string) int {
	dp := make([][]int, N+1) // dp(i, j) := decide i, #greater than last = j
	for i := 0; i <= N; i++ {
		dp[i] = make([]int, N+1)
	}
	for j := 0; j <= N-1; j++ {
		dp[1][j] = 1
	}
	for i := 1; i < N; i++ {
		sum := make([]int, N+2) // sum[j+1] := dp[i][0]+...+dp[i][j]
		for j := 1; j <= N+1; j++ {
			sum[j] = sum[j-1] + dp[i][j-1]
		}
		for j := 0; j <= N; j++ {
			if S[i-1] == '<' {
				dp[i+1][j] += sum[N-i+1] - sum[j+1]
				dp[i+1][j] %= MOD
			} else {
				dp[i+1][j] += sum[j+1]
				dp[i+1][j] %= MOD
			}
		}
	}
	return dp[N][0]
}

func AbsInt(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

func MaxInt(nums ...int) int {
	ret := -INF
	for _, v := range nums {
		if ret < v {
			ret = v
		}
	}
	return ret
}

func MinInt(nums ...int) int {
	ret := math.MaxInt64
	for _, v := range nums {
		if ret > v {
			ret = v
		}
	}
	return ret
}

var _r, _w = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
