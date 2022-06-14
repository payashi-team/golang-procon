package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

const (
	INF = int(1 << 61)
	// MOD = int(1e9 + 7)
	MOD = 998244353
)

func main() {
	var _w, _r = bufio.NewWriter(os.Stdout), bufio.NewReader(os.Stdout)
	defer _w.Flush()
	var N int
	var S []byte
	fmt.Fscan(_r, &N, &S)
	ans := Solve(N, S)
	fmt.Fprintf(_w, "%s\n", ans)
}

type Item struct {
	val, idx int
}

type SegTree struct {
	n     int
	nodes []Item
	f     func(Item, Item) Item
}

func NewSegTree(data []int) *SegTree {
	st := new(SegTree)
	st.n = 1
	for st.n < len(data) {
		st.n *= 2
	}
	st.nodes = make([]Item, st.n*2-1)
	for i := st.n - 1; i < st.n*2-1; i++ {
		if i-(st.n-1) < len(data) {
			st.nodes[i] = Item{data[i-(st.n-1)], i - (st.n - 1)}
		} else {
			st.nodes[i] = Item{INF, -1}
		}
	}
	st.f = func(a, b Item) Item {
		if a.val < b.val {
			return a
		} else if a.val > b.val {
			return b
		} else {
			if a.idx < b.idx {
				return b
			} else {
				return a
			}
		}
	}
	for i := st.n - 2; i >= 0; i-- {
		l, r := i*2+1, i*2+2
		st.nodes[i] = st.f(st.nodes[l], st.nodes[r])
	}
	return st
}

func (st *SegTree) Query(l, r int) Item {
	var dfs func(int, int, int) Item
	dfs = func(k, lb, ub int) Item {
		if l <= lb && ub <= r {
			return st.nodes[k]
		} else if ub <= l || r <= lb {
			return Item{INF, -1}
		} else {
			mid := (lb + ub) / 2
			lv := dfs(2*k+1, lb, mid)
			rv := dfs(2*k+2, mid, ub)
			return st.f(lv, rv)
		}
	}
	return dfs(0, 0, st.n)
}

func Solve(N int, S []byte) string {
	data := make([]int, N)
	for i := 0; i < N; i++ {
		data[i] = int(S[i] - 'a')
	}
	T := make([]byte, N)
	for i := 0; i < N; i++ {
		T[i] = S[i]
	}
	st := NewSegTree(data)
	l, r := 0, N
	for l+1 < r {
		item := st.Query(l+1, r)
		// fmt.Printf("l: %d, r: %d, nextr: %d\n", l, r, item.idx)
		if item.val != INF && item.val < data[l] {
			// swap l, item.idx
			T[l], T[item.idx] = T[item.idx], T[l]
			r = item.idx
		}
		l++
	}
	return string(T)
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

// func nextInt() int {
// 	_s.Scan()
// 	i, e := strconv.Atoi(_s.Text())
// 	if e != nil {
// 		panic(e)
// 	}
// 	return i
// }

// func nextLine() string {
// 	_s.Scan()
// 	return _s.Text()
// }
