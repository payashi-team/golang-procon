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

type Edge struct {
	x, y int
}

func main() {
	defer _w.Flush()
	var N int
	fmt.Fscan(_r, &N)
	E := make([][]int, N)
	for i := 0; i < N-1; i++ {
		var x, y int
		fmt.Fscan(_r, &x, &y)
		x--
		y--
		E[x] = append(E[x], y)
		E[y] = append(E[y], x)
	}
	ans := Solve(N, E)
	fmt.Fprintf(_w, "%d\n", ans)
}

func Solve(N int, E [][]int) int {
	used := make([]bool, N)
	memo := make(map[[2]int]int)
	var dfs func(int, int) int // c: 0=white, 1=black
	dfs = func(v, c int) int {
		if memo[[2]int{v, c}] > 0 {
			return memo[[2]int{v, c}]
		}
		// なんでおなじ引数で複数回よばれてる？？木構造なのに
		// 色のバリエーションがあるからや！せやメモ化しよ
		// paint in color c
		ret := 1
		for _, u := range E[v] {
			if used[u] {
				continue
			}
			used[u] = true
			if c == 0 {
				ret *= dfs(u, 0) + dfs(u, 1)
			} else {
				ret *= dfs(u, 0)
			}
			ret %= MOD
			used[u] = false
		}
		memo[[2]int{v, c}] = ret
		return ret
	}
	used[0] = true
	return (dfs(0, 0) + dfs(0, 1)) % MOD
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
