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
	H := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(_r, &H[i])
	}
	ans := Solve(N, K, H)
	fmt.Fprintf(_w, "%d\n", ans)
}

func Solve(N, K int, H []int) int {
	dp := make([]int, N)
	for i := 0; i < N; i++ {
		dp[i] = INF
	}
	dp[0] = 0
	for i := 0; i < N; i++ {
		for j := 1; j <= MinInt(i, K); j++ {
			dp[i] = MinInt(dp[i], dp[i-j]+AbsInt(H[i]-H[i-j]))
		}
	}
	return dp[N-1]
}

func AbsInt(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

func MaxInt(nums ...int) int {
	ret := -1
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
