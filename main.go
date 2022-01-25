package main

import (
	"bufio"
	"fmt"
	"os"
)

// const (
// MOD = int(1e9 + 7)
// MOD = 998244353
// )

func main() {
	defer _w.Flush()
	var N int
	fmt.Fscan(_r, &N)
	A := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(_r, &A[i])
	}
	ans := Solve(N, A)
	fmt.Fprintf(_w, "%d\n", ans)
}

func Solve(N int, A []int) int {
	ret := 0
	avg := 0
	for _, v := range A {
		avg += v
	}
	if avg%N != 0 {
		return -1
	}
	avg /= N
	sum := 0
	start := -1
	for i := 0; i < N; i++ {
		if start < 0 {
			if A[i] != avg {
				start = i
				sum = A[i]
			}
		} else {
			sum += A[i]
			if sum == avg*(i-start+1) {
				ret += i - start
				start = -1
			}
		}
	}
	return ret
}

var _r, _w = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
