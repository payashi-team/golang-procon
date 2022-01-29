package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	INF = int(1 << 62)
	// MOD = int(1e9 + 7)
	MOD = 998244353
)

func main() {
	defer _w.Flush()
	var N int
	fmt.Fscan(_r, &N)
	T := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(_r, &T[i])
	}
	ans := Solve(N, T)
	fmt.Fprintf(_w, "%d\n", ans)
}

func Solve(N int, T []int) int {
	S := 0
	for _, v := range T {
		S += v
	}
	dp := make([][]int, N+1)
	for i := 0; i <= N; i++ {
		dp[i] = make([]int, S/2+1)
	}
	for i := 0; i < N; i++ {
		for j := 0; j <= S/2; j++ {
			if j-T[i] >= 0 {
				dp[i+1][j] = MaxInt(dp[i][j], dp[i][j-T[i]]+T[i])
			} else {
				dp[i+1][j] = dp[i][j]
			}
		}
	}
	return S-dp[N][S/2]
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

var _r, _w = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
