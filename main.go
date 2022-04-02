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
	A := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(_r, &A[i])
	}
	ans := Solve(N, A)
	fmt.Fprintf(_w, "%d\n", ans)
}

func Solve(N int, A []int) int {
	dp := make([][]int, N+1) // dp(i, j):=[i, j) mincost
	for i := 0; i <= N; i++ {
		dp[i] = make([]int, N+1)
		if i == N {
			continue
		}
	}
	S := make([]int, N+1) // S[i+1] = A[0]+A[1]+...+A[i]
	for i := 0; i < N; i++ {
		S[i+1] = S[i] + A[i]
	}
	var rec func(int, int) int
	rec = func(l, r int) int {
		if l+1 == r {
			return 0
		}
		if dp[l][r] > 0 {
			return dp[l][r]
		}
		ret := INF
		for m := l + 1; m < r; m++ {
			ret = MinInt(ret, rec(l, m)+rec(m, r)+(S[r]-S[l]))
		}
		dp[l][r] = ret
		return ret
	}
	return rec(0, N)
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
