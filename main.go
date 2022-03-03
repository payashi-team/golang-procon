package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

const (
	INF = int(1 << 61)
	MOD = int(1e9 + 7)
	// MOD = 998244353
)

func main() {
	defer _w.Flush()
	var N, K int
	fmt.Fscan(_r, &N, &K)
	A := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(_r, &A[i])
	}
	ans := Solve(N, K, A)
	fmt.Fprintf(_w, "%d\n", ans)
}

type SegTree struct {
	n     int
	nodes []int
}

func NewSegTree(n_ int) *SegTree {
	st := new(SegTree)
	n := 1
	for n < n_ {
		n *= 2
	}
	st.n = n
	st.nodes = make([]int, 2*n-1)
	return st
}

func (st *SegTree) Update(pos, x int) {
	pos += st.n - 1
	st.nodes[pos] = x
	for pos > 0 {
		par := (pos - 1) / 2
		sib := 4*par + 3 - pos
		st.nodes[par] = MaxInt(st.nodes[pos], st.nodes[sib])
		pos = par
	}
}

func (st *SegTree) Query(a, b int) int {
	if a < 0 {
		a = 0
	}
	if b > st.n {
		b = st.n
	}
	var query func(int, int, int) int
	query = func(k, l, r int) int {
		if a <= l && r <= b {
			return st.nodes[k]
		} else if r <= a || b <= l {
			return 0
		} else {
			mid := (l + r) / 2
			lv := query(2*k+1, l, mid)
			rv := query(2*k+2, mid, r)
			return MaxInt(lv, rv)
		}
	}
	return query(0, 0, st.n)
}

func Solve(N, K int, A []int) int {
	MAX_A := int(3e5)
	st := NewSegTree(MAX_A + 1)
	for i := 0; i < N; i++ {
		max := st.Query(-K+A[i], K+A[i]+1) + 1
		st.Update(A[i], max)
	}
	return st.nodes[0]
}

func AbsInt(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

func MaxInt(nums ...int) int {
	ret := -1
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
