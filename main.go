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
	l, r := 0.5, float64(1e9)+1
	for r-l > 1e-5 {
		mid := (l + r) / 2
		cnt := 0
		for i := 0; i < N; i++ {
			cnt += int(math.Ceil(float64(A[i])/mid)) - 1
		}
		if cnt > K {
			l = mid
		} else {
			r = mid
		}
	}
	return int(math.Ceil(math.Round(r*(1e4)) / (1e4)))
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
