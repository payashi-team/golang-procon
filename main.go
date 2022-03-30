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
	A := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(_r, &A[i])
	}
	ans := Solve(N, K, A)
	if ans {
		fmt.Fprintf(_w, "First\n")
	} else {
		fmt.Fprintf(_w, "Second\n")
	}
}

func Solve(N, K int, A []int) bool {
	dp := make([]bool, K+1) // true: First, false: Second
	for i := 0; i <= K; i++ {
		sum := false
		for _, v := range A {
			if i-v < 0 {
				continue
			}
			if !dp[i-v] {
				sum = true
				break
			}
		}
		dp[i] = sum
	}
	return dp[K]
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
