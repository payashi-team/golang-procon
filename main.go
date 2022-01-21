package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	// MOD = int(1e9 + 7)
	MOD = 998244353
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
	ans := Solve(N, A, B)
	fmt.Fprintf(_w, "%d\n", ans)
}

func Solve(N int, A, B []int) int {
	C := make([]int, 3001)
	for i := A[0]; i <= B[0]; i++ {
		C[i] = 1
	}
	for i := 1; i < N; i++ {
		sum := 0
		for j := 0; j <= 3000; j++ {
			sum += C[j]
			sum %= MOD
			if j > B[i] || j < A[i] {
				C[j] = 0
			} else {
				C[j] = sum
			}
		}
	}
	ret := 0
	for i := 0; i <= 3000; i++ {
		ret += C[i]
		ret %= MOD
	}
	return ret
}

var _r, _w = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
