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
	var N, M int
	fmt.Fscan(_r, &N, &M)
	edges := make([][]int, N)
	for i := 0; i < M; i++ {
		var x, y int
		fmt.Fscan(_r, &x, &y)
		x--
		y--
		edges[x] = append(edges[x], y)
	}
	ans := Solve(N, M, edges)
	fmt.Fprintf(_w, "%d\n", ans)
}

func Solve(N, M int, edges [][]int) int {
	que := make([]int, 0)
	outs := make([]int, N)
	ins := make([]int, N)
	for i := 0; i < N; i++ {
		for _, v := range edges[i] {
			ins[v]++
		}
		outs[i] = len(edges[i])
	}
	for i := 0; i < N; i++ {
		if ins[i] == 0 {
			que = append(que, i)
		}
	}
	type Edge struct {
		form, to int
	}
	topo := make([]int, 0)
	used := make(map[Edge]bool)
	for len(que) > 0 {
		u := que[0]
		que = que[1:]
		topo = append(topo, u)
		for _, v := range edges[u] {
			if used[Edge{u, v}] {
				continue
			}
			used[Edge{u, v}] = true
			ins[v]--
			if ins[v] == 0 {
				que = append(que, v)
			}
		}
	}
	dp := make([]int, N)
	for _, u := range topo {
		for _, v := range edges[u] {
			dp[v] = MaxInt(dp[v], dp[u]+1)
		}
	}
	return MaxInt(dp...)
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
