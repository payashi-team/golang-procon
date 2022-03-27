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
	ps := make([]float64, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(_r, &ps[i])
	}
	ans := Solve(N, ps)
	fmt.Fprintf(_w, "%.12f\n", ans)
}

func Solve(N int, ps []float64) float64 {
	dp := make([][]float64, N+1)
	for i := 0; i <= N; i++ {
		dp[i] = make([]float64, N+1)
	}
	dp[0][0] = 1.0
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if dp[i][j] == 0 {
				continue
			}
			dp[i+1][j] += dp[i][j] * (1 - ps[i])
			dp[i+1][j+1] += dp[i][j] * ps[i]
		}
	}
	ans := 0.0
	for i := N/2 + 1; i <= N; i++ {
		ans += dp[N][i]
	}
	return ans
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
