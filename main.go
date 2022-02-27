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
	fmt.Fprintf(_w, "%d\n", ans)
}

func Solve(N, K int, A []int) int {
	sum := make([]int, N+1)
	pos := 0
	used := make([]int, N)
	for i := 0; i < N; i++ {
		used[i] = -1
	}
	for t := 0; t < K; t++ {
		if used[pos] >= 0 {
			// the end of the loop
			p := used[pos]
			l := t - p
			lcost := sum[t] - sum[p]
			a, b := (K-p)/l, (K-p)%l
			return sum[p] + lcost*a + (sum[p+b] - sum[p])
		}
		used[pos] = t
		sum[t+1] = sum[t] + A[pos]
		pos = sum[t+1] % N
	}
	return sum[K]
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
