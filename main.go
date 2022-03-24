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
	dp := make([]map[int]int, N+1) // key: v, value: w
	for i := 0; i <= N; i++ {
		dp[i] = make(map[int]int)
	}
	dp[0][0] = 0
	for i := 0; i < N; i++ {
		item := items[i]
		for j := 0; j <= N*int(1e3); j++ {
			w1, ok1 := dp[i][j]
			if !ok1 {
				w1 = INF
			}
			w2, ok2 := dp[i][j-item.v]
			if !ok2 {
				w2 = INF
			} else {
				w2 += item.w
			}
			if w := MinInt(w1, w2); w != INF {
				dp[i+1][j] = w
			}
		}
	}
	ret := 0
	for v, w := range dp[N] {
		if w <= W {
			ret = MaxInt(ret, v)
		}
	}
	return ret
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
