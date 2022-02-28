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
	var N, Q int
	fmt.Fscan(_r, &N, &Q)
	X := make([]int, N)
	R := make([]int, N)
	H := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(_r, &X[i], &R[i], &H[i])
	}
	A := make([]int, Q)
	B := make([]int, Q)
	for i := 0; i < Q; i++ {
		fmt.Fscan(_r, &A[i], &B[i])
	}
	ans := Solve(N, Q, X, R, H, A, B)
	for _, v := range ans {
		fmt.Fprintf(_w, "%.8f\n", v)
	}
}

func Solve(N, Q int, X, R, H, A, B []int) []float64 {
	volume := make([]float64, int(2e4)+1)
	for i := 0; i < N; i++ {
		v := math.Pi * float64(R[i]*R[i]) / (float64(H[i]) * float64(H[i]) * 3.)
		for j := 0; j < H[i]; j++ {
			h := float64(H[i] - j)
			volume[X[i]+j] += v * ((h * h * h) - ((h - 1) * (h - 1) * (h - 1)))
		}
	}
	for i := 1; i <= int(2e4); i++ {
		volume[i] += volume[i-1]
	}
	ans := make([]float64, Q)
	for i := 0; i < Q; i++ {
		if A[i] == 0 {
			ans[i] = volume[B[i]-1]
		} else {
			ans[i] = volume[B[i]-1] - volume[A[i]-1]
		}
	}
	return ans
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
