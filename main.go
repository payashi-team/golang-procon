package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

const (
	INF = int(1 << 61)
	// MOD = int(1e9 + 7)
	MOD = 998244353
)

func main() {
	defer _w.Flush()
	var N, X, Y int
	fmt.Fscan(_r, &N, &X, &Y)
	A := make([]int, N)
	B := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(_r, &A[i], &B[i])
	}
	ans := Solve(N, X, Y, A, B)
	fmt.Fprintf(_w, "%d\n", ans)
}

func Solve(N, X, Y int, A, B []int) int {
	dp := make([][][]int, N+1)
	// dp(i, x, y):=sumA=x, sumB=y, min #bento
	for i := 0; i <= N; i++ {
		dp[i] = make([][]int, X+1)
		for j := 0; j <= X; j++ {
			dp[i][j] = make([]int, Y+1)
			for k := 0; k <= Y; k++ {
				dp[i][j][k] = INF
			}
		}
	}
	dp[0][0][0] = 0
	// dp(i+1, j, k) = dp(i, j, k)
	// dp(i+1, j+Ai, k+Bi) = dp(i, j, k)+1
	for i := 0; i < N; i++ {
		for j := 0; j <= X; j++ {
			for k := 0; k <= Y; k++ {
				dp[i+1][j][k] = MinInt(dp[i+1][j][k], dp[i][j][k])
				dp[i+1][MinInt(j+A[i], X)][MinInt(k+B[i], Y)] = MinInt(dp[i+1][MinInt(j+A[i], X)][MinInt(k+B[i], Y)], dp[i][j][k]+1)
			}
		}
	}
	ans := dp[N][X][Y]
	if ans == INF {
		return -1
	} else {
		return ans
	}

}

func Contains(x int, nums ...int) bool {
	for _, v := range nums {
		if v == x {
			return true
		}
	}
	return false
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
