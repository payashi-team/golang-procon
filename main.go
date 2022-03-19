package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

const (
	INF = int(1 << 61)
	MOD = int(1e9 + 7)
	// MOD = 998244353
)

func main() {
	defer _w.Flush()
	var N string
	var K int
	fmt.Fscan(_r, &N, &K)
	ans := Solve(N, K)
	fmt.Fprintf(_w, "%d\n", ans)
}

func Solve(N string, K int) int {
	L := len(N)
	dp := make([][][2]int, L+1)
	for i := 0; i <= L; i++ {
		dp[i] = make([][2]int, K+1)
	}
	dp[0][0][0] = 1
	for i := 0; i < L; i++ {
		v, _ := strconv.Atoi(N[i : i+1])
		for j := 0; j <= K; j++ {
			for k := 0; k < 2; k++ {
				for d := 0; d < 10; d++ {
					var ni, nj, nk = i + 1, j, k
					if d != 0 {
						nj++
					}
					if nj > K {
						continue
					}
					if k == 0 {
						if d > v {
							continue
						} else if d < v {
							nk = 1
						}
					}
					dp[ni][nj][nk] += dp[i][j][k]
				}
			}
		}
	}
	return dp[L][K][0] + dp[L][K][1]
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
