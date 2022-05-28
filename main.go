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
	N, M, Q := ScanInt(), ScanInt(), ScanInt()
	row := make([]int, N)
	for i := 0; i < N; i++ {
		row[i] = -1
	}
	st := NewSegTree(M + 1)
	for q := 0; q < Q; q++ {
		cmd := ScanInt()
		switch cmd {
		case 1:
			l, r, x := ScanInt(), ScanInt(), ScanInt()
			st.Add(l-1, x)
			st.Add(r, -x)
		case 2:
			i, x := ScanInt(), ScanInt()
			row[i-1] += x
		case 3:
			i, j := ScanInt(), ScanInt()
			ret := 0
			ret += row[i-1]
			ret += st.Query(0, j)
			fmt.Fprintf(_w, "%d\n", ret)
		}
		fmt.Fprintf(_w, "%v\n", st.nodes)
		fmt.Fprintf(_w, "%v\n", row)
	}
}

type SegTree struct {
	n     int
	nodes []int
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

func (st *SegTree) Add(pos, x int) {
	pos += st.n - 1
	for pos > 0 {
		par := (pos - 1) / 2
		st.nodes[pos] += x
		pos = par
	}
	st.nodes[0] += x
}

func (st *SegTree) Query(l, r int) int {
	var dfs func(int, int, int) int
	dfs = func(lb, ub, k int) int {
		if l <= lb && ub <= r {
			return st.nodes[k]
		} else if ub <= l || r <= lb {
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
