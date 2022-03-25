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
	var S, T string
	fmt.Fscan(_r, &S, &T)
	ans := Solve(S, T)
	fmt.Fprintf(_w, "%s\n", ans)
}

func Solve(S, T string) string {
	ls := len(S)
	lt := len(T)
	dp := make([][]int, ls+1)
	for i := 0; i <= ls; i++ {
		dp[i] = make([]int, lt+1)
	}
	for i := 0; i < ls; i++ {
		for j := 0; j < lt; j++ {
			if S[i] == T[j] {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				s1 := dp[i][j+1]
				s2 := dp[i+1][j]
				dp[i+1][j+1] = MaxInt(s1, s2)
			}
		}
	}
	ret := make([]byte, dp[ls][lt])
	for i, j := ls, lt; dp[i][j] > 0; {
		if S[i-1] == T[j-1] {
			ret[dp[i][j]-1] = S[i-1]
			i--
			j--
		} else if dp[i][j] == dp[i-1][j] {
			i--
		} else {
			j--
		}
	}
	return string(ret)
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
