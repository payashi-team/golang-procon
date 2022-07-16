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

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	N, M := ni(), ni()
	edges := make([][]int, N)
	for i := 0; i < M; i++ {
		a, b := ni()-1, ni()-1
		edges[a] = append(edges[a], b)
	}
	Solve(N, M, edges)
}

func Solve(N, M int, edges [][]int) {
	cnt := 0
	uf := NewUfind(N)
	ret := make([]int, N)
	ret[N-1] = 0
	for u := N - 1; u > 0; u-- {
		cnt++
		for _, v := range edges[u] {
			if uf.Unite(u, v) {
				cnt--
			}
		}
		ret[u-1] = cnt
	}
	for i := 0; i < N; i++ {
		fmt.Fprintf(wr, "%d\n", ret[i])
	}
}

type UFind struct {
	par, dep []int
}

func NewUfind(N int) *UFind {
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

func (uf *UFind) Unite(x, y int) bool {
	x = uf.Root(x)
	y = uf.Root(y)

	if uf.Same(x, y) {
		return false
	}
	if uf.dep[x] < uf.dep[y] {
		x, y = y, x
	} else if uf.dep[x] == uf.dep[y] {
		uf.dep[x]++
	}
	uf.par[y] = x
	return true
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
