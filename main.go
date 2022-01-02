package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	defer _w.Flush()
	var H, W int
	fmt.Fscan(_r, &H, &W)
	grid := make([]byte, H*W)
	for i := 0; i < H; i++ {
		var S string
		fmt.Fscan(_r, &S)
		for j := 0; j < W; j++ {
			grid[i*W+j] = S[j]
		}
	}
	ans := Solve(H, W, grid)
	fmt.Fprintf(_w, "%d\n", ans)
}

func Solve(H, W int, grid []byte) int {
	dp:=make([]int, H*W)
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			dp[i*W+j] = int(2e4)
		}
	}
	if grid[0]=='#'{
		dp[0] = 1
	}else{
		dp[0] = 0
	}
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			cur:=i*W+j
			if i>0{
				top:=(i-1)*W+j
				if grid[cur]=='#'&&grid[top]=='.'{
					dp[cur] = MinInt(dp[cur], dp[top]+1)
				}else{
					dp[cur] = MinInt(dp[cur], dp[top])
				}
			}
			if j>0{
				left:=i*W+j-1
				if grid[cur]=='#'&&grid[left]=='.'{
					dp[cur] = MinInt(dp[cur], dp[left]+1)
				}else{
					dp[cur] = MinInt(dp[cur], dp[left])
				}
			}
		}
	}
	return dp[H*W-1]
}

func MinInt(nums ...int)int{
	ret:=int(2e4)
	for _, v := range nums {
		if ret>v{
			ret = v
		}
	}
	return ret
}

var _r, _w = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
