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

type Item struct {
	v, w int
}

func main() {
	defer _w.Flush()
	var N, W int
	fmt.Fscan(_r, &N, &W)
	items := make([]Item, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(_r, &items[i].w, &items[i].v)
	}
	ans := Solve(N, W, items)
	fmt.Fprintf(_w, "%d\n", ans)
}

func Solve(N, W int, items []Item) int {
	dp := make([][]int, N+1)
	for i := 0; i <= N; i++ {
		dp[i] = make([]int, W+1)
	}
	for i := 0; i < N; i++ {
		item := items[i]
		for j := 0; j <= W; j++ {
			if j-item.w >= 0 {
				dp[i+1][j] = MaxInt(dp[i][j], dp[i][j-item.w]+item.v)
			} else {
				dp[i+1][j] = dp[i][j]
			}
		}
	}
	return dp[N][W]
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
