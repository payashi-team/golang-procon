package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	INF = int(1 << 60)

// MOD = int(1e9 + 7)
// MOD = 998244353
)

func main() {
	defer _w.Flush()
	var N, M int
	fmt.Fscan(_r, &N, &M)
	A := make([]int, M)
	B := make([]int, M)
	for i := 0; i < M; i++ {
		fmt.Fscan(_r, &A[i], &B[i])
	}
	ans := Solve(N, M, A, B)
	for _, v := range ans {
		fmt.Fprintf(_w, "%d\n", v)
	}
}

func Solve(N, M int, A, B []int) []int {
	dist := make([][]int, N)
	for i := 0; i < N; i++ {
		dist[i] = make([]int, N)
		for j := 0; j < N; j++ {
			dist[i][j] = INF
		}
		dist[i][i] = 0
	}
	for i := 0; i < M; i++ {
		a := A[i] - 1
		b := B[i] - 1
		dist[a][b] = 1
		dist[b][a] = 1
	}
	for k := 0; k < N; k++ {
		for i := 0; i < N; i++ {
			for j := 0; j < N; j++ {
				dist[i][j] = MinInt(dist[i][j], dist[i][k]+dist[k][j])
			}
		}
	}
	ret := make([]int, N)
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if dist[i][j] == 2 {
				ret[i]++
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
