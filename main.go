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
	var N, A, B int
	fmt.Fscan(_r, &N, &A, &B)
	S := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(_r, &S[i])
	}
	P, Q := Solve(N, A, B, S)
	if P < 0 {
		fmt.Fprintf(_w, "-1\n")
	} else {
		fmt.Fprintf(_w, "%.8f %.8f\n", P, Q)
	}
}

func Solve(N, A, B int, S []int) (float64, float64) {
	A0 := 0.
	max := -1
	min := int(1e9)
	for _, v := range S {
		if max < v {
			max = v
		}
		if min > v {
			min = v
		}
		A0 += float64(v)
	}
	A0 /= float64(N)
	B0 := float64(max - min)
	if B0 == 0 {
		return -1, -1
	}
	P := float64(B) / B0
	Q := float64(A) - P*A0
	return P, Q
}

var _r, _w = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
