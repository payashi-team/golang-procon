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
	H := make([][3]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(_r, &H[i][0], &H[i][1], &H[i][2])
	}
	ans := Solve(N, H)
	fmt.Fprintf(_w, "%d\n", ans)
}

func Solve(N int, H [][3]int) int {
	dp := make([][3]int, N+1)
	for i := 0; i < N; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				if j == k {
					continue
				}
				dp[i+1][j] = MaxInt(dp[i+1][j], dp[i][k]+H[i][j])
			}
		}
	}
	return MaxInt(dp[N][0], dp[N][1], dp[N][2])
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
