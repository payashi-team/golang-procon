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

type Edge struct {
	a, b int
}

func main() {
	defer _w.Flush()
	var N int
	fmt.Fscan(_r, &N)
	M, E := Solve(N)
	fmt.Fprintf(_w, "%d\n", M)
	for _, e := range E {
		fmt.Fprintf(_w, "%d %d\n", e.a, e.b)
	}
}

func Solve(N int) (int, []Edge) {
	odd := N&1 == 1
	if odd {
		N--
	}
	M := N * (N - 2) / 2
	E := make([]Edge, M)
	cur := 0
	for i := 1; i <= N; i++ {
		for j := i + 1; j <= N; j++ {
			if i+j == N+1 {
				continue
			}
			E[cur] = Edge{i, j}
			cur++
		}
	}
	if odd {
		M += N
		for i := 1; i <= N; i++ {
			E = append(E, Edge{i, N + 1})
		}
	}
	return M, E
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
