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
	ans := Solve(N)
	fmt.Fprintf(_w, "%d\n", ans)
}

func Solve(N int) int {
	n := float64(N)
	rt := FloorInt(math.Sqrt(n))
	ret := 0
	for i := 1; i <= FloorInt(n/float64(rt+1)); i++ {
		ret += N / i
	}
	for i := 1; i <= rt; i++ {
		cnt := FloorInt(n/float64(i)) - FloorInt(n/float64(i+1))
		ret += i * cnt
	}
	return ret
}

func FloorInt(x float64) int {
	return int(math.Floor(x))
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
