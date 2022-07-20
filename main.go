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

type Point struct {
	x, y, idx int
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	N := ni()
	ps := make([]Point, N)
	for i := 0; i < N; i++ {
		ps[i] = Point{ni(), ni(), i}
	}
	ans := Solve(N, ps)
	fmt.Fprintf(wr, "%d\n", ans)
}

type Edge struct {
	u, v, cost int
}

func Solve(N int, ps []Point) int {
	sort.Slice(ps, func(i, j int) bool { return ps[i].x < ps[j].x })
	edges := make([]Edge, 0)
	for i := 0; i < N-1; i++ {
		edges = append(edges, Edge{ps[i].idx, ps[i+1].idx, ps[i+1].x - ps[i].x})
	}
	sort.Slice(ps, func(i, j int) bool { return ps[i].y < ps[j].y })
	for i := 0; i < N-1; i++ {
		edges = append(edges, Edge{ps[i].idx, ps[i+1].idx, ps[i+1].y - ps[i].y})
	}
	sort.Slice(edges, func(i, j int) bool { return edges[i].cost < edges[j].cost })
	cost := 0
	uf := NewUFind(N)
	cnt := 0
	for _, e := range edges {
		if uf.Same(e.u, e.v) {
			continue
		}
		uf.Unite(e.u, e.v)
		cost += e.cost
		cnt++
		if cnt == N-1 {
			break
		}
	}
	return cost
}

type UFind struct {
	dep, par []int
}

func NewUFind(N int) *UFind {
	uf := new(UFind)
	uf.par = make([]int, N)
	uf.dep = make([]int, N)
	for i := 0; i < N; i++ {
		uf.par[i] = i
	}
	return uf
}

func (uf *UFind) Root(x int) int {
	if uf.par[x] == x {
		return x
	}
	uf.par[x] = uf.Root(uf.par[x])
	return uf.par[x]
}

func (uf *UFind) Same(x, y int) bool {
	return uf.Root(x) == uf.Root(y)
}

func (uf *UFind) Unite(x, y int) {
	x = uf.Root(x)
	y = uf.Root(y)
	if x == y {
		return
	}
	if uf.dep[x] == uf.dep[y] {
		uf.dep[x]++
	} else if uf.dep[x] < uf.dep[y] {
		x, y = y, x
	}
	uf.par[y] = x
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
