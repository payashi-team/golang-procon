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
	field := make([]string, H)
	for i := 0; i < H; i++ {
		fmt.Fscan(_r, &field[i])
	}
	ans := Solve(H, W, field)
	fmt.Fprintf(_w, "%d\n", ans)
}

func Solve(H, W int, field []string) int {
	dp := make([][]int, H)
	r := make([][]int, H)
	d := make([][]int, H)
	rd := make([][]int, H)
	for i := 0; i < H; i++ {
		dp[i] = make([]int, W)
		r[i] = make([]int, W)
		d[i] = make([]int, W)
		rd[i] = make([]int, W)
	}
	dp[H-1][W-1] = 1
	for h := H - 1; h >= 0; h-- {
		for w := W - 1; w >= 0; w-- {
			if (h == H-1 && w == W-1) || field[h][w] == '#' {
				continue
			}
			if w+1 < W && field[h][w+1] != '#' {
				r[h][w] = r[h][w+1] + dp[h][w+1]
				r[h][w] %= MOD
			}
			if h+1 < H && field[h+1][w] != '#' {
				d[h][w] = d[h+1][w] + dp[h+1][w]
				d[h][w] %= MOD
			}
			if w+1 < W && h+1 < H && field[h+1][w+1] != '#' {
				rd[h][w] = rd[h+1][w+1] + dp[h+1][w+1]
				rd[h][w] %= MOD
			}
			dp[h][w] = r[h][w] + d[h][w] + rd[h][w]
			dp[h][w] %= MOD
		}
	}
	return dp[0][0]
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
