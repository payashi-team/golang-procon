package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	INF = int(1 << 60)
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
		var a, b, c int
		fmt.Fscan(_r, &a, &b, &c)
		dist[a-1][b-1] = c
	}
	ans := Solve(N, M, dist)
	fmt.Fprintf(_w, "%d\n", ans)
}

func Solve(N, M int, dist [][]int) int {
	ret := 0
	for k := 0; k < N; k++ {
		for i := 0; i < N; i++ {
			for j := 0; j < N; j++ {
				dist[i][j] = MinInt(dist[i][j], dist[i][k]+dist[k][j])
				if dist[i][j] != INF {
					ret += dist[i][j]
				}
			}
		}
	}
	return ret
}

func MinInt(nums ...int) int {
	ret := INF
	for _, v := range nums {
		if ret > v {
			ret = v
		}
	}
	return ret
}

var _r, _w = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
