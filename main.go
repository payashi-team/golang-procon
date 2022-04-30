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
	l := 0
	r := int(2e9) + 1
	for r-l > 1 {
		// [l, r)
		mid := (r + l) / 2
		cnt := 0
		for _, v := range A {
			if v <= mid {
				continue
			}
			cnt += v - mid
		}
		if cnt <= K {
			r = mid
		} else {
			l = mid
		}
	}
	ret := 0
	for _, v := range A {
		if v <= r {
			continue
		}
		ret += v*(v+1)/2 - r*(r+1)/2
		K -= v - r
	}
	if r > 0 {
		for _, v := range A {
			if v >= r {
				K--
				if K < 0 {
					break
				}
				ret += r
			}
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
