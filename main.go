package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	INF = int(1 << 60)
	// MOD = int(1e9 + 7)
	MOD = 998244353
)

func main() {
	defer _w.Flush()
	var N, K int
	fmt.Fscan(_r, &N, &K)
	L := make([]int, K)
	R := make([]int, K)
	for i := 0; i < K; i++ {
		fmt.Fscan(_r, &L[i], &R[i])
	}
	ans := Solve(N, K, L, R)
	fmt.Fprintf(_w, "%d\n", ans)
}

func Solve(N, K int, L, R []int) int {
	dp := make([]int, N)
	sdp := make([]int, N+1)
	dp[N-1] = 1
	sdp[1] = dp[N-1]
	getSdp := func(i int) int {
		if N-i < 0 {
			return sdp[0]
		}
		return sdp[N-i]
	}
	for i := N - 2; i >= 0; i-- {
		for k := 0; k < K; k++ {
			l, r := L[k], R[k]
			dp[i] += (getSdp(i+l) - getSdp(i+r+1) + MOD) % MOD
			dp[i] %= MOD
		}
		sdp[N-i] = sdp[N-i-1] + dp[i]
		sdp[N-i] %= MOD
	}
	return dp[0]
}

var _r, _w = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
