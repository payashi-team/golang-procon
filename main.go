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
	B := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(_r, &A[i])
	}
	for i := 0; i < N; i++ {
		fmt.Fscan(_r, &B[i])
	}
	a, b := Solve(N, A, B)
	fmt.Fprintf(_w, "%d\n%d\n", a, b)
}

func Solve(N int, A, B []int) (int, int) {
	mp := make(map[int]int)
	for i := 0; i < N; i++ {
		mp[A[i]] = i + 1
	}
	var a, b int
	for i := 0; i < N; i++ {
		v := mp[B[i]]
		if v == 0 {
			continue
		}
		if v == i+1 {
			a++
		} else {
			b++
		}
	}
	return a, b
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
