package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

const (
	INF = int(1 << 61)
	// MOD = int(1e9 + 7)
	MOD = 998244353
)

var sc = bufio.NewScanner(os.Stdin)
var wr = bufio.NewWriter(os.Stdout)

type Edge struct {
	from, to, cost, qi int
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	N := ni()
	M := ni()
	Q := ni()
	edges := make([]Edge, M+Q)
	for i := 0; i < M; i++ {
		a, b, c := ni()-1, ni()-1, ni()
		if a == b {
			continue
		}
		if a > b {
			a, b = b, a
		}
		edges[i] = Edge{a, b, c, -1}
	}
	ans := make([]bool, Q)
	for q := 0; q < Q; q++ {
		u, v, w := ni()-1, ni()-1, ni()
		if u > v {
			u, v = v, u
		}
		edges[M+q] = Edge{u, v, w, q}
	}
	uf := NewUnionFind(N)
	sort.Slice(edges, func(i, j int) bool { return edges[i].cost < edges[j].cost })
	for _, e := range edges {
		if e.qi < 0 {
			if !uf.Same(e.from, e.to) {
				uf.Unite(e.from, e.to)
			}
		} else {
			ans[e.qi] = !uf.Same(e.from, e.to)
		}
	}
	for _, v := range ans {
		if v {
			fmt.Fprintln(wr, "Yes")
		} else {
			fmt.Fprintln(wr, "No")
		}
	}
}

type UnionFind struct {
	par, dep []int
}

func NewUnionFind(N int) *UnionFind {
	uf := new(UnionFind)
	uf.par = make([]int, N)
	uf.dep = make([]int, N)
	for i := 0; i < N; i++ {
		uf.par[i] = i
	}
	return uf
}

func (uf *UnionFind) Root(x int) int {
	if uf.par[x] == x {
		return x
	}
	uf.par[x] = uf.Root(uf.par[x])
	return uf.par[x]
}

func (uf *UnionFind) Same(x, y int) bool {
	return uf.Root(x) == uf.Root(y)
}

func (uf *UnionFind) Unite(x, y int) {
	x = uf.Root(x)
	y = uf.Root(y)
	if x == y {
		return
	}
	if uf.dep[x] > uf.dep[y] {
		uf.par[y] = x
	} else if uf.dep[x] == uf.dep[y] {
		uf.par[x] = y
		uf.dep[y]++
	} else {
		uf.par[x] = y
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

func ni() int {
	sc.Scan()
	x, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return x
}

func nl() string {
	sc.Scan()
	return sc.Text()
}
