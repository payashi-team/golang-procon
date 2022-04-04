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
	fmt.Fscan(_r, &N)
	A := make([][]int, N)
	for i := 0; i < N; i++ {
		A[i] = make([]int, N)
		for j := 0; j < N; j++ {
			fmt.Fscan(_r, &A[i][j])
		}
	}
	ans := Solve(N, A)
	fmt.Fprintf(_w, "%d\n", ans)
}

func Solve(N int, A [][]int) int {
	dp := make([][]int, N+1) // dp[i+1][j] := (1<<i)&~(1<<j)
	for i := 0; i <= N; i++ {
		dp[i] = make([]int, 1<<N)
	}
	dp[0][0] = 1
	for i := 0; i < N; i++ {
		for bit := 1<<(i+1) - 1; bit < 1<<N; {
			for j := 0; j < N; j++ {
				if A[i][j] == 1 && (bit>>j)&1 == 1 {
					dp[i+1][bit] += dp[i][bit&^(1<<j)]
					dp[i+1][bit] %= MOD
				}
			}
			x := bit & -bit
			y := bit + x
			bit = (((bit & ^y) / x) >> 1) | y
		}
	}
	ret := 0
	for i := 0; i < N; i++ {
		ret += dp[N][(1<<N-1)&^(1<<i)]
	}
	return dp[N][(1<<N)-1]
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
