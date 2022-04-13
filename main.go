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
	var N, K int
	fmt.Fscan(_r, &N, &K)
	A := make([][]int, N)
	for i := 0; i < N; i++ {
		A[i] = make([]int, N)
		for j := 0; j < N; j++ {
			fmt.Fscan(_r, &A[i][j])
		}
	}
	ans := Solve(N, K, A)
	fmt.Fprintf(_w, "%d\n", ans)
}

type UnionFind struct {
	parent []int
	depth  []int
	size   []int
}

func NewUnionFind(N int) *UnionFind {
	uf := new(UnionFind)
	uf.parent = make([]int, N)
	uf.depth = make([]int, N)
	uf.size = make([]int, N)
	for i := 0; i < N; i++ {
		uf.parent[i] = i
		uf.size[i] = 1
	}
	return uf
}

func (uf *UnionFind) Root(x int) int {
	if uf.parent[x] == x {
		return x
	}
	uf.parent[x] = uf.Root(uf.parent[x])
	return uf.parent[x]
}

func (uf *UnionFind) Same(x, y int) bool {
	return uf.Root(x) == uf.Root(y)
}

func (uf *UnionFind) Unite(x, y int) {
	if uf.Same(x, y) {
		return
	}
	x = uf.Root(x)
	y = uf.Root(y)
	if uf.depth[x] > uf.depth[y] {
		uf.parent[y] = x
		uf.size[x] += uf.size[y]
	} else {
		uf.parent[x] = y
		uf.size[y] += uf.size[x]
		if uf.depth[x] == uf.depth[y] {
			uf.depth[y]++
		}
	}
}

func Solve(N, K int, A [][]int) int {
	rows := NewUnionFind(N)
	cols := NewUnionFind(N)
	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			ok_rows := true
			ok_cols := true
			for k := 0; k < N; k++ {
				if A[i][k]+A[j][k] > K {
					ok_rows = false
				}
				if A[k][i]+A[k][j] > K {
					ok_cols = false
				}
			}
			if ok_rows {
				rows.Unite(i, j)
			}
			if ok_cols {
				cols.Unite(i, j)
			}
		}
	}
	fact := make([]int, N+1)
	fact[0] = 1
	for i := 0; i < N; i++ {
		fact[i+1] = fact[i] * (i + 1)
		fact[i+1] %= MOD
	}
	ret := 1
	used := make([]bool, N)
	for i := 0; i < N; i++ {
		r := rows.Root(i)
		if used[r] {
			continue
		}
		used[r] = true
		ret *= fact[rows.size[r]]
		ret %= MOD
	}
	used = make([]bool, N)
	for i := 0; i < N; i++ {
		r := cols.Root(i)
		if used[r] {
			continue
		}
		used[r] = true
		ret *= fact[cols.size[r]]
		ret %= MOD
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
