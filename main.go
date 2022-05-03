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
	B := make([]int, N-1)
	for i := 0; i < N-1; i++ {
		if A[i+1] > A[i] {
			B[i] = 1
		}
	}
	ret := 0
	cnt := 0
	B = append(B, 0)
	for i := N - 2; i >= 0; i-- {
		if B[i] == 0 {
			ret += MaxInt(0, cnt-(K-1)+1)
			cnt = 0
		} else {
			cnt++
			B[i] = cnt
		}
	}
	ret += MaxInt(0, cnt-(K-1)+1)
	return ret
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
