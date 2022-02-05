package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const (
	// INF = int(1 << 60)
	// MOD = int(1e9 + 7)
	MOD = 998244353
)

func main() {
	defer _w.Flush()
	var N, Q int
	fmt.Fscan(_r, &N, &Q)
	L := make([]int, Q)
	R := make([]int, Q)
	for i := 0; i < Q; i++ {
		fmt.Fscan(_r, &L[i], &R[i])
	}
	ans := Solve(N, Q, L, R)
	if ans {
		fmt.Fprintf(_w, "Yes\n")
	} else {
		fmt.Fprintf(_w, "No\n")
	}
}

type Range struct {
	l, r int
}

func Solve(N, Q int, L, R []int) bool {
	qs := make([]Range, Q)
	for i := 0; i < Q; i++ {
		qs[i] = Range{L[i] - 1, R[i] - 1}
	}
	sort.Slice(qs, func(i, j int) bool {
		if qs[i].r == qs[j].r {
			return qs[i].l < qs[j].l
		} else {
			return qs[i].r < qs[j].r
		}
	})
	dp := make([]bool, N+1)
	dp[0] = true
	pos := 0
	for i := 0; i < N; i++ {
		// fmt.Printf("pos: %d\n", pos)
		for pos < Q && qs[pos].r == i {
			if dp[qs[pos].l] {
				dp[i] = true
			}
			if dp[i] {
				dp[qs[pos].l] = true
			}
			pos++
		}
	}
	// fmt.Printf("%v\n", dp)
	return dp[N-1]
}

var _r, _w = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
