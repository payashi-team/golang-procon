package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

const (
	INF = int(1 << 61)
	// MOD = int(1e9 + 7)
	MOD = 998244353
)

func main() {
	defer _w.Flush()
	var N int
	fmt.Fscanf(_r, "%d\n", &N)
	A := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(_r, &A[i])
	}
	ans := Solve(N, A)
	if ans {
		fmt.Fprintf(_w, "Win\n")
	} else {
		fmt.Fprintf(_w, "Lose\n")
	}
}

func Solve(N int, A []int) bool {
	if N&1 == 1 {
		return true
	}
	xor := 0
	for _, v := range A {
		xor ^= v

	}
	for _, v := range A {
		if xor == v {
			return true
		}
	}
	return false
}

// Both a and b are sorted
func MinDist(a, b []int) int {
	ret := INF
	for _, v := range a {
		idx := sort.Search(len(b), func(i int) bool { return v <= b[i] })
		if idx == 0 {
			ret = MinInt(ret, b[idx]-v)
		} else if idx == len(b) {
			ret = MinInt(ret, v-b[idx-1])
		} else {
			ret = MinInt(ret, b[idx]-v, v-b[idx-1])
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
