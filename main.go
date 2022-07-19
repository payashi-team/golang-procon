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
	S, T, M := ni(), ni(), ni()
	edges := make([][]int, S+T) // 0<=u<S, S<=v<S+T
	for i := 0; i < M; i++ {
		u, v := ni()-1, ni()-1
		edges[u] = append(edges[u], v-S)
	}
	Solve(S, T, M, edges)
}

func Solve(S, T, M int, edges [][]int) {
	cnt := make([][]int, T)
	for i := 0; i < T; i++ {
		cnt[i] = make([]int, T)
		for j := 0; j < T; j++ {
			cnt[i][j] = -1
		}
	}
	for u := 0; u < S; u++ {
		for _, v1 := range edges[u] {
			for _, v2 := range edges[u] {
				if v1 == v2 {
					continue
				} else if cnt[v1][v2] < 0 {
					cnt[v1][v2] = u
				} else {
					fmt.Fprintf(wr, "%d %d %d %d\n", v1+1+S, v2+1+S, u+1, cnt[v1][v2]+1)
					return
				}
			}
		}
	}
	fmt.Fprintf(wr, "-1\n")
}

type SegTree struct {
	nodes []Item
	n     int
}

type Item struct {
	idx, val int
}

func NewSegTree(n int) *SegTree {
	st := new(SegTree)
	st.n = 1
	for st.n < n {
		st.n *= 2
	}
	st.nodes = make([]Item, st.n*2-1)
	for i := 1; i <= st.n; i++ {
		if i <= n {
			st.nodes[i-1+st.n-1] = Item{i, i}
		} else {
			st.nodes[i-1+st.n-1] = Item{-1, INF}
		}
	}
	for i := st.n - 2; i >= 0; i-- {
		l, r := i*2+1, i*2+2
		if st.nodes[l].val < st.nodes[r].val {
			st.nodes[i] = st.nodes[l]
		}
	}
	return st
}

// idx is 1-indexed
func (st *SegTree) Update(x, idx int) {
	pos := idx - 1 + st.n - 1
	for pos > 0 {
		st.nodes[pos].val = x
		par := (pos - 1) / 2
		sib := 4*par + 3 - pos
		x = MinInt(st.nodes[pos].val, st.nodes[sib].val)
		pos = par
	}
	st.nodes[0].val = x
}

func (st *SegTree) Query(l, r int) Item {
	var dfs func(int, int, int) Item
	dfs = func(lb, ub, k int) Item {
		if l <= lb && ub <= r {
			return st.nodes[k]
		} else if r <= lb || ub <= l {
			return Item{-1, INF}
		} else {
			mid := (lb + ub) / 2
			lv := dfs(lb, mid, 2*k+1)
			rv := dfs(mid, ub, 2*k+2)
			if lv.val > rv.val {
				return rv
			} else {
				return lv
			}
		}
	}
	return dfs(0, st.n, 0)
}

type UFind struct {
	par, dep, size []int
}

func NewUfind(N int) *UFind {
	uf := new(UFind)
	uf.par = make([]int, N)
	uf.dep = make([]int, N)
	uf.size = make([]int, N)
	for i := 0; i < N; i++ {
		uf.par[i] = i
		uf.size[i] = 1
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

func (uf *UFind) Size(x int) int {
	x = uf.Root(x)
	return uf.size[x]
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
	uf.size[x] += uf.size[y]
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
