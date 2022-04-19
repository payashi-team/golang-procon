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
	ans := Solve(N)
	fmt.Fprintf(_w, "%d\n", len(ans))
	for _, v := range ans {
		fmt.Fprintf(_w, "%d\n", v)
	}
}

func Solve(N int) []int {
	ret := make([]int, 0)
	for i := MaxInt(N-9*18, 1); i <= N; i++ {
		dsum := 0
		num := i
		for num > 0 {
			dsum += num % 10
			num /= 10
		}
		if dsum+i == N {
			ret = append(ret, i)
		}
	}
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
