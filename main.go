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
	fmt.Fprintf(_w, "%d\n", Solve(N))
}

func Solve(N int) int {

	f := func(a, b int) int {
		return a*a*a + a*a*b + a*b*b + b*b*b
	}
	j := int(1e6)
	X := INF
	for i := 0; j >= 0 && i <= int(1e6); i++ {
		for ; j >= 0 && f(i, j) >= N; j-- {
			X = MinInt(X, f(i, j))
		}
	}
	return X
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
