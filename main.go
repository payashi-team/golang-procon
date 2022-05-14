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
	A := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(_r, &A[i])
	}
	ans := Solve(N, A)
	fmt.Fprintf(_w, "%d\n", ans)
}

func Solve(N int, A []int) int {
	calc := func(i, j int) (int, int) {
		takahashi := 0
		aoki := 0
		if i > j {
			i, j = j, i
		}
		for k := i; k <= j; k++ {
			if (k-i)%2 == 0 {
				takahashi += A[k]
			} else {
				aoki += A[k]
			}
		}
		return takahashi, aoki
	}
	ret := -INF
	// takahashi's pos
	for i := 0; i < N; i++ {
		aoki := -INF
		maxj := -1
		// aoki's pos
		for j := 0; j < N; j++ {
			if i == j {
				continue
			}
			_, score := calc(i, j)
			if score > aoki {
				aoki = score
				maxj = j
			}
		}
		takahashi, _ := calc(i, maxj)
		ret = MaxInt(ret, takahashi)
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
