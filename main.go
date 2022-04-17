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
	var N, M, K, S, T, X int
	fmt.Fscan(_r, &N, &M, &K, &S, &T, &X)
	U := make([]int, M)
	V := make([]int, M)
	for i := 0; i < M; i++ {
		fmt.Fscan(_r, &U[i], &V[i])
	}
	ans := Solve(N, M, K, S, T, X, U, V)
	fmt.Fprintf(_w, "%d\n", ans)
}

func Solve(N, M, K, S, T, X int, U, V []int) int {
	S--
	T--
	X--
	edges := make([][]int, N)
	for i := 0; i < N; i++ {
		edges[i] = make([]int, 0)
	}
	for i := 0; i < M; i++ {
		u := U[i] - 1
		v := V[i] - 1
		edges[u] = append(edges[u], v)
		edges[v] = append(edges[v], u)
	}
	dp := make([][][]int, K+1)
	for i := 0; i <= K; i++ {
		dp[i] = make([][]int, N)
		for j := 0; j < N; j++ {
			dp[i][j] = make([]int, 2)
		}
	}
	// dp(i, j, b) := S -> j (#path = i),
	// b = 0 (when passing through X even times), 1 (otherwise)
	dp[0][S][0] = 1
	for i := 0; i < K; i++ {
		for j := 0; j < N; j++ {
			for b := 0; b <= 1; b++ {
				for _, k := range edges[j] {
					// S -> j -> k
					if k == X {
						dp[i+1][k][(b+1)%2] += dp[i][j][b]
						dp[i+1][k][(b+1)%2] %= MOD
					} else {
						dp[i+1][k][b] += dp[i][j][b]
						dp[i+1][k][b] %= MOD
					}
				}
			}
		}
	}
	return dp[K][T][0]
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
