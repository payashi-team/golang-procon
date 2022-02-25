package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

const (
	INF = int(1 << 61)
	MOD = int(1e9 + 7)
	// MOD = 998244353
)

func main() {
	defer _w.Flush()
	var N int
	fmt.Fscan(_r, &N)
	A := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(_r, &A[i])
	}
	M, ops := Solve(N, A)
	fmt.Fprintf(_w, "%d\n", M)
	for _, op := range ops {
		fmt.Fprintf(_w, "%d %d\n", op.x, op.y)
	}
}

type Op struct {
	x, y int
}

func Solve(N int, A []int) (int, []Op) {
	sort.Ints(A)
	m := sort.SearchInts(A, 0) - 1
	if m == N-1 {
		m--
	} else if m == -1 {
		m++
	}
	ret := 0
	ops := make([]Op, N-1)
	for i := m + 1; i < N-1; i++ {
		ops[i-(m+1)] = Op{A[m], A[i]}
		A[m] -= A[i]
	}
	for i := 0; i <= m; i++ {
		ops[i+N-2-m] = Op{A[N-1], A[i]}
		A[N-1] -= A[i]
	}
	ret = A[N-1]
	return ret, ops
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
