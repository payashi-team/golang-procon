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
	var N, M int
	fmt.Fscan(_r, &N, &M)
	dist := make([][]int, N)
	for i := 0; i < N; i++ {
		dist[i] = make([]int, N)
		for j := 0; j < N; j++ {
			dist[i][j] = INF
		}
		dist[i][i] = 0
	}
	for i := 0; i < M; i++ {
		var a, b, t int
		fmt.Fscan(_r, &a, &b, &t)
		a--
		b--
		dist[a][b] = t
		dist[b][a] = t
	}
	ans := Solve(N, M, dist)
	fmt.Fprintf(_w, "%d\n", ans)
}

func Solve(N, M int, dist [][]int) int {
	for k := 0; k < N; k++ {
		for i := 0; i < N; i++ {
			for j := 0; j < N; j++ {
				dist[i][j] = MinInt(dist[i][j], dist[i][k]+dist[k][j])
			}
		}
	}
	ret := INF
	for i := 0; i < N; i++ {
		ret = MinInt(ret, MaxInt(dist[i]...))
	}
	return ret
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
