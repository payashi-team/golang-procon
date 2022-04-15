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
	var A string
	fmt.Fscan(_r, &A)
	ans := Solve(A)
	fmt.Fprintf(_w, "%d\n", ans)
}

func Solve(A string) int {
	asym := 0
	N := len(A)
	for i := 0; i < N-1-i; i++ {
		if A[i] != A[N-1-i] {
			asym++
		}
	}
	if asym == 0 {
		return (N - N&1) * 25
	} else if asym == 1 {
		return N*25 - 2
	}else{
		return N*25
	}
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
