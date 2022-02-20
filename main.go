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

type Point struct {
	x, y, z int
}

func main() {
	defer _w.Flush()
	var N int
	fmt.Fscan(_r, &N)
	var S string
	fmt.Fscan(_r, &S)
	T := make([]string, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(_r, &T[i])
	}
	ans := Solve(N, S, T)
	fmt.Fprintf(_w, "%d\n", ans)
}

func Solve(N int, S string, T []string) int {
	dp := make([]int, len(S)+1)
	dp[0] = 1
	for i := 1; i <= len(S); i++ {
		for j := 0; j < N; j++ {
			d := len(T[j])
			if i-d >= 0 && S[i-d:i] == T[j] {
				dp[i] += dp[i-d]
				dp[i] %= MOD
			}
		}
	}
	return dp[len(S)]
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
