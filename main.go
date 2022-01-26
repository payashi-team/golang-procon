package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// const (
// MOD = int(1e9 + 7)
// MOD = 998244353
// )

func main() {
	defer _w.Flush()
	var T, N, M int
	fmt.Fscan(_r, &T, &N)
	A := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(_r, &A[i])
	}
	fmt.Fscan(_r, &M)
	B := make([]int, M)
	for i := 0; i < M; i++ {
		fmt.Fscan(_r, &B[i])
	}
	ans := Solve(T, N, M, A, B)
	if ans {
		fmt.Fprintf(_w, "yes\n")
	} else {
		fmt.Fprintf(_w, "no\n")
	}
}

func Solve(T, N, M int, A, B []int) bool {
	if N < M {
		return false
	}
	sort.Ints(A)
	sort.Ints(B)
	j := 0
	for i := 0; i < M; i++ {
		for ; j < N && (A[j] > B[i] || B[i] > A[j]+T); j++ {
		}
		if j == N {
			return false
		}
		j++
	}
	return true
}

var _r, _w = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
