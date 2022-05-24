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
	uf := NewUnionFind(N)
	for i := 0; i < M; i++ {
		var a, b int
		fmt.Fscan(_r, &a, &b)
		a--
		b--
		uf.Unite(a, b)
	}
	mp := make(map[int]struct{})
	for i := 0; i < N; i++ {
		mp[uf.Root(i)] = struct{}{}
	}
	fmt.Fprintf(_w, "%d\n", len(mp)-1)
}

type UnionFind struct {
	n             int
	parent, depth []int
}

func NewUnionFind(n int) *UnionFind {
	uf := new(UnionFind)
	uf.n = n
	uf.parent = make([]int, n)
	uf.depth = make([]int, n)
	for i := 0; i < n; i++ {
		uf.parent[i] = i
	}
	return uf
}

func (uf *UnionFind) Root(x int) int {
	if x == uf.parent[x] {
		return x
	}
	uf.parent[x] = uf.Root(uf.parent[x])
	return uf.parent[x]
}

func (uf *UnionFind) Unite(x, y int) {
	x = uf.Root(x)
	y = uf.Root(y)

	if x == y {
		return
	}
	if uf.depth[x] < uf.depth[y] {
		uf.parent[x] = y
	} else {
		uf.parent[y] = x
		if uf.depth[x] == uf.depth[y] {
			uf.depth[x]++
		}
	}
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
