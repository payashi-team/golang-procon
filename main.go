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
	var N, M int
	fmt.Fscan(_r, &N, &M)
	A := make([]int, M)
	for i := 0; i < M; i++ {
		fmt.Fscan(_r, &A[i])
	}
	ans := Solve(N, M, A)
	for _, v := range ans {
		fmt.Fprintf(_w, "%d\n", v)
	}
}

func Solve(N, M int, A []int) []int {
	used := make(map[int]bool)
	ret := make([]int, N)
	pos := 0
	for i := M - 1; i >= 0; i-- {
		a := A[i]
		if used[a] {
			continue
		}
		used[A[i]] = true
		ret[pos] = A[i]
		pos++
	}
	for i := 1; i <= N; i++ {
		if used[i] {
			continue
		}
		ret[pos] = i
		pos++
	}
	return ret
}

var _r, _w = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
