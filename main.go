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
	var N int
	fmt.Fscan(_r, &N)
	D := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(_r, &D[i])
	}
	max, min := Solve(N, D)
	fmt.Fprintf(_w, "%d\n%d\n", max, min)
}

func Solve(N int, D []int) (int, int) {
	sum := 0
	e := MaxInt(D...)
	for _, v := range D {
		sum += v
	}
	if e <= sum-e {
		return sum, 0
	} else {
		return sum, 2*e - sum
	}
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

