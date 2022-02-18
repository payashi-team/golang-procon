package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	INF = int(1 << 60)
	// MOD = int(1e9 + 7)
	MOD = 998244353
)

func main() {
	defer _w.Flush()
	var N int
	fmt.Fscan(_r, &N)
	A := make([]int, N+1)
	for i := 0; i <= N; i++ {
		fmt.Fscan(_r, &A[i])
	}
	ans := Solve(N, A)
	fmt.Fprintf(_w, "%d\n", ans)
}

type Range struct {
	min, max int
}

func Solve(N int, A []int) int {
	ret := make([]Range, N+1)
	ret[N] = Range{A[N], A[N]}
	for i := N - 1; i >= 0; i-- {
		ret[i].min = (ret[i+1].min+1)/2 + A[i]
		ret[i].max = ret[i+1].max + A[i]
	}
	ret[0].max = 1
	if ret[0].max < 1 || 1 < ret[0].min {
		return -1
	}
	for i := 1; i <= N; i++ {
		ret[i].max = MinInt((ret[i-1].max-A[i-1])*2, ret[i].max)
	}
	sum := 0
	for i := 0; i <= N; i++ {
		sum += ret[i].max
	}
	return sum
}

func MinInt(nums ...int) int {
	ret := INF
	for _, v := range nums {
		if ret > v {
			ret = v
		}
	}
	return ret
}

var _r, _w = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
