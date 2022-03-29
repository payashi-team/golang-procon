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
	A := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(_r, &A[i])
	}
	ans := Solve(N, A)
	fmt.Fprintf(_w, "%.12f\n", ans)
}

func Solve(N int, A []int) float64 {
	type State struct {
		a, b, c int
	}
	var init State
	for _, v := range A {
		if v == 1 {
			init.a++
		} else if v == 2 {
			init.b++
		} else if v == 3 {
			init.c++
		}
	}
	n := float64(N)
	var dp [301][301][301]float64
	for c := 0; c <= N; c++ {
		for b := 0; b <= N; b++ {
			for a := 0; a <= N; a++ {
				if a+b+c == 0 {
					continue
				}
				val := 1.0
				if a > 0 {
					val += dp[a-1][b][c] * float64(a) / n
				}
				if b > 0 {
					val += dp[a+1][b-1][c] * float64(b) / n
				}
				if c > 0 {
					val += dp[a][b+1][c-1] * float64(c) / n
				}
				val *= n / float64(a+b+c)
				dp[a][b][c] = val
			}
		}
	}
	return dp[init.a][init.b][init.c]
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
