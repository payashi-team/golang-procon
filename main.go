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

func main() {
	defer _w.Flush()
	_s.Split(bufio.ScanWords)
	_s.Buffer([]byte{}, math.MaxInt32)
	N, Q := ScanInt(), ScanInt()
	A := make([]int, N)
	B := make([]int, N)
	for i := 0; i < N; i++ {
		A[i] = ScanInt()
	}
	for i := 0; i < N; i++ {
		B[i] = ScanInt()
	}
	diffA := make([]int, N-1)
	diffB := make([]int, N-1)
	for i := 0; i < N-1; i++ {
		diffA[i] = AbsInt(A[i+1] - A[i])
		diffB[i] = AbsInt(B[i+1] - B[i])
	}
	sta := NewSegTree(N-1, diffA)
	stb := NewSegTree(N-1, diffB)
	for q := 0; q < Q; q++ {
		h1, h2, w1, w2 := ScanInt()-1, ScanInt()-1, ScanInt()-1, ScanInt()-1
		gcd := GCD(sta.Query(h1, h2), stb.Query(w1, w2))
		gcd = GCD(gcd, A[h1]+B[w1])
		fmt.Fprintf(_w, "%d\n", gcd)
	}
}

type SegTree struct {
	n     int
	nodes []int
}

func NewSegTree(n int, nums []int) *SegTree {
	st := new(SegTree)
	st.n = 1
	for st.n < n {
		st.n *= 2
	}
	st.nodes = make([]int, st.n*2-1)
	for i := 0; i < len(nums); i++ {
		st.nodes[i+st.n-1] = nums[i]
	}
	for i := st.n - 2; i >= 0; i-- {
		lv, rv := st.nodes[i*2+1], st.nodes[i*2+2]
		st.nodes[i] = GCD(lv, rv)
	}
	return st
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
			return GCD(lv, rv)
		}
	}
	return dfs(0, st.n, 0)
}

func GCD(x, y int) int {
	if x > y {
		x, y = y, x
	}
	if x == 0 {
		return y
	}
	return GCD(y%x, x)
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

func ScanInt() int {
	_s.Scan()
	num, err := strconv.Atoi(_s.Text())
	if err != nil {
		panic(err)
	}
	return num
}

var _s, _w = bufio.NewScanner(os.Stdin), bufio.NewWriter(os.Stdout)
