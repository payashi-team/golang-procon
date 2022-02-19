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
	S := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(_r, &S[i])
	}
	ans := Solve(N, K, S)
	fmt.Fprintf(_w, "%d\n", ans)
}

func Solve(N, K int, S []int) int {
	min := INF
	for _, v := range S {
		min = MinInt(min, v)
	}
	if min == 0 {
		return N
	} else if K == 0 {
		return 0
	}
	r := 0
	cur := 1
	ret := 0
	for l := 0; l < N; l++ {
		for r < N && cur*S[r] <= K {
			cur *= S[r]
			r++
		}
		// fmt.Printf("[%d, %d): %d\n", l, r, cur)
		ret = -MinInt(-ret, -(r - l))
		if l == r {
			r++
		} else {
			cur /= S[l]
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
