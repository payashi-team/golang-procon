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
	var N, M int
	fmt.Fscan(_r, &N, &M)
	A := make([]int, N)
	B := make([]int, M)
	for i := 0; i < N; i++ {
		fmt.Fscan(_r, &A[i])
	}
	for i := 0; i < M; i++ {
		fmt.Fscan(_r, &B[i])
	}
	ans := Solve(N, M, A, B)
	if ans {
		fmt.Fprintf(_w, "Yes\n")
	} else {
		fmt.Fprintf(_w, "No\n")
	}
}

func Solve(N, M int, A, B []int) bool {
	cnta := make(map[int]int)
	cntb := make(map[int]int)
	for _, v := range A {
		cnta[v]++
	}
	for _, v := range B {
		cntb[v]++
	}
	for k, v := range cntb {
		if cnta[k] < v {
			return false
		}
	}
	return true
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
