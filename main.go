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

	var N, M int
	fmt.Fscan(_r, &N, &M)
	B := make([][]int, N)
	for i := 0; i < N; i++ {
		B[i] = make([]int, M)
		var S string
		fmt.Fscan(_r, &S)
		for j := 0; j < M; j++ {
			B[i][j] = int(S[j] - '0')
		}
	}
	ans := Solve(N, M, B)
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			fmt.Fprintf(_w, "%d", ans[i][j])
		}
		fmt.Fprintln(_w)
	}
}

func Solve(N, M int, B [][]int) [][]int {
	A := make([][]int, N)
	for i := 0; i < N; i++ {
		A[i] = make([]int, M)
	}
	for i := 1; i < N-1; i++ {
		for j := 1; j < M-1; j++ {
			if i-2 >= 0 {
				A[i][j] = B[i-1][j] - A[i-2][j] - A[i-1][j-1] - A[i-1][j+1]
			} else {
				A[i][j] = B[i-1][j]
			}
		}
	}
	return A
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
