package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	defer _w.Flush()
	var N, M int
	fmt.Fscan(_r, &N, &M)
	A := make([]int, M)
	B := make([]int, M)
	for i := 0; i < M; i++ {
		fmt.Fscan(_r, &A[i], &B[i])
	}
	ans := Solve(N, M, A, B)
	if ans {
		fmt.Fprintf(_w, "YES\n")
	} else {
		fmt.Fprintf(_w, "NO\n")
	}
}

func Solve(N, M int, A, B []int) bool {
	cnt := make(map[int]int)
	for i := 0; i < M; i++ {
		cnt[A[i]]++
		cnt[B[i]]++
	}
	for _, v := range cnt {
		if v%2 == 1 {
			return false
		}
	}
	return true
}

var _r, _w = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
