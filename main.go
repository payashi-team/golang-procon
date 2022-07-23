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
	N := ni()
	C := make([]int, N)
	X := make([]int, N)
	for i := 0; i < N; i++ {
		C[i] = ni()
	}
	for i := 0; i < N; i++ {
		X[i] = ni()
	}
	ans := Solve(N, C, X)
	fmt.Fprintf(wr, "%d\n", ans)
}

func Solve(N int, C, X []int) int {
	mp := make(map[int][]int)
	for i := 0; i < N; i++ {
		mp[C[i]] = append(mp[C[i]], X[i])
	}
	ret := CountInv(N, X)
	// for _, v := range mp {
	// 	ret -= CountInv(N, v)
	// }
	return ret
}

func CountInv(N int, A []int) int {
	st := NewSegTree(N + 1)
	ret := 0
	for i := 0; i < len(A); i++ {
		st.Add(1, A[i])
		fmt.Fprintf(wr, "(0, %d], %d\n", A[i], st.Query(0, A[i]+1))
		ret += i + 1 - st.Query(0, A[i]+1)
	}
	return ret
}

type SegTree struct {
	nodes []int
	n     int
}

func NewSegTree(n int) *SegTree {
	st := new(SegTree)
	st.n = 1
	for st.n < n {
		st.n *= 2
	}
	st.nodes = make([]int, st.n*2-1)
	return st
}

func (st *SegTree) Add(x, idx int) {
	pos := idx + st.n - 1
	for pos > 0 {
		st.nodes[pos] += x
		par := (pos - 1) / 2
		sib := 4*par + 3 - pos
		x = st.nodes[pos] + st.nodes[sib]
		pos = par
	}
	st.nodes[0] += x
}

func (st *SegTree) Query(l, r int) int {
	var dfs func(int, int, int) int
	dfs = func(lb, ub, k int) int {
		if l <= lb && ub <= r {
			return st.nodes[k]
		} else if r <= lb || ub <= l {
			return 0
		} else {
			mid := (lb + ub) / 2
			lv := dfs(lb, mid, 2*k+1)
			rv := dfs(mid, ub, 2*k+2)
			return lv + rv
		}
	}
	return dfs(0, st.n, 0)
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
