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
	var N int
	fmt.Fscan(_r, &N)
	P := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(_r, &P[i])
	}
	ans := Solve(N, P)
	fmt.Fprintf(_w, "%d\n", ans)
}

func Solve(N int, P []int) int {
	Q := make([]int, N)
	for i := 0; i < N; i++ {
		Q[P[i]-1] = i
	}
	ret := -1
	cnt := 1
	for i := 1; i < N; i++ {
		if Q[i] > Q[i-1] {
			cnt++
		} else {
			ret = MaxInt(ret, cnt)
			cnt = 1
		}
	}
	ret = MaxInt(ret, cnt)
	ret = N - ret
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
