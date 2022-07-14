package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
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
	x, y int
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	N := ni()
	s := Point{ni(), ni()}
	t := Point{ni(), ni()}
	ps := make([]Point, N)
	R := make([]int, N)
	for i := 0; i < N; i++ {
		ps[i] = Point{ni(), ni()}
		R[i] = ni()
	}
	Solve(N, s, t, ps, R)
}

func Solve(N int, s, t Point, ps []Point, R []int) {
	uf := NewUnionFind(N + 2)
	ps = append(ps, s, t)
	R = append(R, 0, 0)
	for i := 0; i < N+2; i++ {
		p := ps[i]
		for j := i + 1; j < N+2; j++ {
			q := ps[j]
			dist := (p.x-q.x)*(p.x-q.x) + (p.y-q.y)*(p.y-q.y)
			rmax := (R[i] + R[j]) * (R[i] + R[j])
			rmin := (R[i] - R[j]) * (R[i] - R[j])
			if rmin <= dist && dist <= rmax {
				uf.Unite(i, j)
			}
		}
	}
	if uf.Same(N, N+1) {
		fmt.Fprintln(wr, "Yes")
	} else {
		fmt.Fprintln(wr, "No")
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
