package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const (
	INF = int(1 << 62)

// MOD = int(1e9 + 7)
// MOD = 998244353
)

func main() {
	defer _w.Flush()
	var N, M int
	fmt.Fscan(_r, &N, &M)
	A := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(_r, &A[i])
	}
	X := make([]int, M)
	Y := make([]int, M)
	for i := 0; i < M; i++ {
		fmt.Fscan(_r, &X[i], &Y[i])
	}
	ans := Solve(N, M, A, X, Y)
	fmt.Fprintf(_w, "%d\n", ans)
}

func Solve(N, M int, A, X, Y []int) int {
	edges := make([][]int, N)
	for i := 0; i < M; i++ {
		x := X[i] - 1
		y := Y[i] - 1
		edges[x] = append(edges[x], y)
	}
	type Rank struct {
		index, price int
	}
	cheap := make([]Rank, N)
	for i := 0; i < N; i++ {
		cheap[i] = Rank{i, A[i]}
	}
	sort.Slice(cheap, func(i, j int) bool { return cheap[i].price < cheap[j].price })
	used := make([]bool, N)
	ret := -INF
	for _, p := range cheap {
		que := make([]int, 0)
		que = append(que, p.index)
		for len(que) > 0 {
			u := que[0]
			que = que[1:]
			for _, v := range edges[u] {
				if used[v] {
					continue
				} else {
					used[v] = true
					que = append(que, v)
					ret = MaxInt(ret, A[v]-p.price)
				}
			}
		}
	}
	return ret
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

var _r, _w = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
