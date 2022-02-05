package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	// INF = int(1 << 60)
	// MOD = int(1e9 + 7)
	MOD = 998244353
)

func main() {
	defer _w.Flush()
	var N, Q int
	fmt.Fscan(_r, &N, &Q)
	L := make([]int, Q)
	R := make([]int, Q)
	for i := 0; i < Q; i++ {
		fmt.Fscan(_r, &L[i], &R[i])
	}
	ans := Solve(N, Q, L, R)
	if ans {
		fmt.Fprintf(_w, "Yes\n")
	} else {
		fmt.Fprintf(_w, "No\n")
	}
}

type UnionFind struct {
	par, dep []int
}

func (uf *UnionFind) Root(x int) int {
	if uf.par[x] == x {
		return x
	}
	uf.par[x] = uf.par[uf.par[x]]
	return uf.Root(uf.par[x])
}

func (uf *UnionFind) Unite(a, b int) {
	a = uf.Root(a)
	b = uf.Root(b)
	if a == b {
		return
	}
	if uf.dep[a] < uf.dep[b] {
		uf.par[a] = b
	} else {
		uf.par[b] = a
		if uf.dep[a] == uf.dep[b] {
			uf.dep[a]++
		}
	}
}

func Solve(N, Q int, L, R []int) bool {
	uf := new(UnionFind)
	uf.par = make([]int, N+1)
	uf.dep = make([]int, N+1)
	for i := 0; i <= N; i++ {
		uf.par[i] = i
	}
	for i := 0; i < Q; i++ {
		uf.Unite(L[i]-1, R[i])
	}
	return uf.Root(0) == uf.Root(N)
}

var _r, _w = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
