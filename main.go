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
	var N, A, B int
	fmt.Fscan(_r, &N, &A, &B)
	ans := Solve(N, A, B)
	if ans {
		fmt.Fprintf(_w, "Takahashi\n")
	} else {
		fmt.Fprintf(_w, "Aoki\n")
	}
}

func Solve(N, A, B int) bool {
	if A == B {
		return N%(A+1) != 0
	} else if A > B {
		return true
	} else {
		return N <= A
	}
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
