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
	var N, K int
	fmt.Fscan(_r, &N, &K)
	A := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(_r, &A[i])
	}
	ans := Solve(N, K, A)
	fmt.Fprintf(_w, "%d\n", ans)
}

func Solve(N, K int, A []int) int {
	dp := make([]int, K+1) // dp[i][j] := children xi, candies xj
	dp[0] = 1
	for i := 0; i < N; i++ {
		cum := make([]int, K+2) // cum[j+1] := dp[i][0]+...+dp[i][j]
		for j := 0; j < K+1; j++ {
			cum[j+1] = cum[j] + dp[j]
		}
		for j := 0; j <= K; j++ {
			dp[j] = cum[j+1] - cum[MaxInt(0, j-A[i])] + MOD
			dp[j] %= MOD
		}
	}
	return dp[K]
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
