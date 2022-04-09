package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

const (
	INF = int(1 << 61)
	MOD = int(1e9 + 7)
	// MOD = 998244353
)

func main() {
	defer _w.Flush()
	var K string
	var D int
	fmt.Fscan(_r, &K, &D)
	ans := Solve(K, D)
	fmt.Fprintf(_w, "%d\n", ans)
}

func Solve(K string, D int) int {
	N := len(K)
	A := make([]int, N)
	for i := 0; i < N; i++ {
		v, _ := strconv.Atoi(string(K[N-1-i]))
		A[i] = v
	}
	dp := make([][][]int, N+1) // dp(i, j, k) := i digits, %=j(mod D), <=K(k=1)
	for i := 0; i <= N; i++ {
		dp[i] = make([][]int, D)
		for j := 0; j < D; j++ {
			dp[i][j] = make([]int, 2)
		}
	}
	dp[0][0][0] = 1
	dp[0][0][1] = 1
	for i := 0; i < N; i++ {
		for j := 0; j < D; j++ {
			for t := 0; t < A[i]; t++ {
				dp[i+1][(j+t)%D][1] += dp[i][j][0]
				dp[i+1][(j+t)%D][1] %= MOD
			}
			dp[i+1][(j+A[i])%D][1] += dp[i][j][1]
			dp[i+1][(j+A[i])%D][1] %= MOD
			for t := 0; t < 10; t++ {
				dp[i+1][(j+t)%D][0] += dp[i][j][0]
				dp[i+1][(j+t)%D][0] %= MOD
			}
		}
	}
	return (dp[N][0][1] - 1 + MOD) % MOD
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
