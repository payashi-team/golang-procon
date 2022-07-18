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
	A := make([]int, N)
	B := make([]int, N)
	for i := 0; i < N; i++ {
		A[i] = ni()
		B[i] = ni()
	}
	Solve(N, M, A, B)
}

func Solve(N, M int, A, B []int) {
	ans := make([]int, M+2)
	C := make([][]int, M+1)
	for i := 0; i < N; i++ {
		C[A[i]] = append(C[A[i]], i)
		C[B[i]] = append(C[B[i]], i)
	}
	cnt := make([]int, N)
	ng := N // A, Bどちらも含んでいない組の数
	for i, j := 1, 1; i <= M; i++ {
		for j <= M && ng != 0 {
			for _, x := range C[j] {
				if cnt[x] == 0 {
					ng--
				}
				cnt[x]++
			}
			j++
		}
		if ng != 0 {
			break
		}
		for _, x := range C[i] {
			if cnt[x] == 1 {
				ng++
			}
			cnt[x]--
		}
		ans[j-i]++
		ans[M+1-i+1]--
	}
	for i := 0; i < M; i++ {
		ans[i+1] += ans[i]
		fmt.Fprintf(wr, "%d ", ans[i+1])
	}
	fmt.Fprintln(wr)
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
