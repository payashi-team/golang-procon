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
	var N, R int
	var S string
	fmt.Fscan(_r, &N, &R, &S)
	ans := Solve(N, R, S)
	fmt.Fprintf(_w, "%d\n", ans)
}

func Solve(N, R int, S string) int {
	M := -1
	for i := N - 1; i >= 0; i-- {
		if S[i] == '.' {
			M = i
			break
		}
	}
	ret := 0
	if M-R+1 > 0 {
		ret = M - R + 1
	} else {
		if M < 0 {
			return 0
		} else {
			return 1
		}
	}
	S = S[:M-R+1]
	for pos := 0; pos < M-R+1; {
		if S[pos] == '.' {
			ret++
			pos += R
		} else {
			pos++
		}
	}
	ret++
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
