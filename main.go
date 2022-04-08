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
	A := make([][]int, N)
	for i := 0; i < N; i++ {
		A[i] = make([]int, N)
		for j := 0; j < N; j++ {
			fmt.Fscan(_r, &A[i][j])
		}
	}
	ans := Solve(N, K, A)
	fmt.Fprintf(_w, "%d\n", ans)
}

func Solve(N, K int, A [][]int) int {
	Multiple := func(X, Y [][]int) [][]int {
		Z := make([][]int, N)
		for i := 0; i < N; i++ {
			Z[i] = make([]int, N)
		}
		for i := 0; i < N; i++ {
			for j := 0; j < N; j++ {
				for k := 0; k < N; k++ {
					Z[i][j] += X[i][k] * Y[k][j]
					Z[i][j] %= MOD
				}
			}
		}
		return Z
	}
	Pow := func(X [][]int, p int) [][]int {
		ret := make([][]int, N)
		for i := 0; i < N; i++ {
			ret[i] = make([]int, N)
			ret[i][i] = 1
		}
		for p > 0 {
			if p&1 == 1 {
				ret = Multiple(ret, X)
			}
			p >>= 1
			X = Multiple(X, X)
		}
		return ret
	}
	B := Pow(A, K)
	ret := 0
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			ret += B[i][j]
			ret %= MOD
		}
	}
	return ret
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

var _r, _w = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
