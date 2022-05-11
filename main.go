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
	fmt.Fscan(_r, &N)
	W := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(_r, &W[i])
		W[i]--
	}
	ans := Solve(N, W)
	fmt.Fprintf(_w, "%d\n", ans)
}

func Solve(N int, W []int) int {
	ends := make([]int, 0) // asc
	for i := 0; i < N; i++ {
		// new end
		if len(ends) == 0 || ends[len(ends)-1] < W[i] {
			ends = append(ends, W[i])
			sort.Ints(ends)
		} else {
			idx := sort.Search(len(ends), func(j int) bool { return ends[j] >= W[i] })
			ends[idx] = W[i]
		}
	}
	return len(ends)
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
