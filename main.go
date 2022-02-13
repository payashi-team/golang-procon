package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const (
	INF = int(1 << 60)
	// MOD = int(1e9 + 7)
	MOD = 998244353
)

func main() {
	defer _w.Flush()
	var K, N, M int
	fmt.Fscan(_r, &K, &N, &M)
	A := make([]int, K)
	for i := 0; i < K; i++ {
		fmt.Fscan(_r, &A[i])
	}
	ans := Solve(K, N, M, A)
	for i := 0; i < K; i++ {
		fmt.Fprintf(_w, "%d ", ans[i].value)
	}
	fmt.Fprintln(_w)
}

type Item struct {
	idx, value, remain int
}

// Ci = Ai * (M/N) (float64)
// max|Bi - Ci|を最小化すればよい
// p <= Ci < q (p, qは整数)
// sum(p) <= M < sum(q)
// すべてのBiにpかqを選んだとき max|Bi - Ci| < 1
// Biにpでもqでもない数を一つでも選んだとき |Bi - Ci| >= 1

func Solve(K, N, M int, A []int) []Item {
	more := M
	ret := make([]Item, K)
	for i := 0; i < K; i++ {
		ret[i] = Item{i, (A[i] * M) / N, (A[i] * M) % N}
		more -= ret[i].value
	}
	sort.Slice(ret, func(i, j int) bool { return ret[i].remain > ret[j].remain })
	for i := 0; i < more; i++ {
		ret[i].value++
	}
	sort.Slice(ret, func(i, j int) bool { return ret[i].idx < ret[j].idx })
	return ret
}

var _r, _w = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
