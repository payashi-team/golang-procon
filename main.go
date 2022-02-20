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

type Point struct {
	x, y, z int
}

func main() {
	defer _w.Flush()
	var N int
	fmt.Fscan(_r, &N)
	P := make([]Point, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(_r, &P[i].x, &P[i].y, &P[i].z)
	}
	ans := Solve(N, P)
	fmt.Fprintf(_w, "%d\n", ans)
}

func Solve(N int, P []Point) int {
	dp := make([][]int, 1<<N)
	for i := 0; i < 1<<N; i++ {
		dp[i] = make([]int, N)
		for j := 0; j < N; j++ {
			dp[i][j] = INF
		}
	}
	dist := make([][]int, N)
	for i := 0; i < N; i++ {
		dist[i] = make([]int, N)
		a := P[i]
		for j := 0; j < N; j++ {
			b := P[j]
			dist[i][j] = AbsInt(a.x-b.x) + AbsInt(a.y-b.y) + MaxInt(0, b.z-a.z)
		}
	}
	dp[1<<N-1][0] = 0
	for i := 1<<N - 2; i >= 0; i-- {
		for u := 0; u < N; u++ {
			for v := 0; v < N; v++ {
				if (i>>u)&1 == 0 {
					dp[i][v] = MinInt(dp[i][v], dp[i|(1<<u)][u]+dist[v][u])
				}
			}
		}
	}
	return dp[0][0]
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
