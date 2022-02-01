package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	INF = int(1 << 62)
	// MOD = int(1e9 + 7)
	MOD = 998244353
)

func main() {
	defer _w.Flush()
	var N int
	fmt.Fscan(_r, &N)
	A := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(_r, &A[i])
	}
	ans := Solve(N, A)
	fmt.Fprintf(_w, "%d\n", ans)
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

func (uf *UnionFind) Same(a, b int) bool {
	return uf.Root(a) == uf.Root(b)
}

func (uf *UnionFind) Unite(a, b int) {
	a = uf.Root(a)
	b = uf.Root(b)
	if uf.Same(a, b) {
		return
	}
	if uf.dep[a] < uf.dep[b] {
		uf.par[a] = b
	} else if uf.dep[a] > uf.dep[b] {
		uf.par[b] = a
	} else {
		uf.par[b] = a
		uf.dep[a]++
	}
}

func Solve(N int, A []int) int {
	uf := new(UnionFind)
	uf.par = make([]int, int(2e5)+1)
	uf.dep = make([]int, int(2e5)+1)
	for i := 0; i <= int(2e5); i++ {
		uf.par[i] = i
	}
	ret := 0
	for i := 0; i < (N+1)/2; i++ {
		a := A[i]
		b := A[N-1-i]
		if !uf.Same(a, b) {
			uf.Unite(a, b)
			ret++
		}
	}
	return ret
}

var _r, _w = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
