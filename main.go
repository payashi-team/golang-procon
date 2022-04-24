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
	defer _w.Flush()
	var T int
	fmt.Fscan(_r, &T)
	for i := 0; i < T; i++ {
		var N, A, B, X, Y, Z int
		fmt.Fscan(_r, &N, &A, &B, &X, &Y, &Z)
		ans := Solve(N, A, B, X, Y, Z)
		fmt.Fprintf(_w, "%d\n", ans)
	}
}

func Solve(N, A, B, X, Y, Z int) int {
	Y = MinInt(Y, X*A)
	Z = MinInt(Z, X*B)
	// A no houga cospa ii
	if A*Z < B*Y {
		A, B = B, A
		Y, Z = Z, Y
	}
	// Aop <= N/A
	// Bop < A
	ans := N * X
	if N/A < A-1 {
		for i := 0; i*A <= N; i++ {
			j := (N - i*A) / B
			k := N - i*A - j*B
			ans = MinInt(ans, i*Y+j*Z+k*X)
		}
	} else {
		for j := 0; j < A && j*B <= N; j++ {
			i := (N - j*B) / A
			k := N - i*A - j*B
			ans = MinInt(ans, i*Y+j*Z+k*X)
		}
	}
	return ans
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

var _r, _w = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
