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
	var H, W int
	fmt.Fscan(_r, &H, &W)
	cells := make([][]bool, H)
	for i := 0; i < H; i++ {
		cells[i] = make([]bool, W)
		var tmp string
		fmt.Fscan(_r, &tmp)
		for j := 0; j < W; j++ {
			if tmp[j] == '.' {
				cells[i][j] = true
			}
		}
	}
	ans := Solve(H, W, cells)
	fmt.Fprintf(_w, "%d\n", ans)
}

func Solve(H, W int, cells [][]bool) int {
	dp := make([][]int, H)
	for i := 0; i < H; i++ {
		dp[i] = make([]int, W)
	}
	dp[0][0] = 1
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			if !cells[i][j] {
				continue
			}
			if i > 0 {
				dp[i][j] += dp[i-1][j]
			}
			if j > 0 {
				dp[i][j] += dp[i][j-1]
			}
			dp[i][j] %= MOD
		}
	}
	return dp[H-1][W-1]
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
