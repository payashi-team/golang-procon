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
	var a, b, c, x int
	fmt.Fscan(_r, &a, &b, &c, &x)
	ans := Solve(a, b, c, x)
	fmt.Fprintf(_w, "%.8f\n", ans)
}

func Solve(a, b, c, x int) float64 {
	if x <= a {
		return 1
	} else if x <= b {
		return float64(c) / float64(b-a)
	} else {
		return 0
	}
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
