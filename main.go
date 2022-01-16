package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	defer _w.Flush()
	var N int
	fmt.Fscan(_r, &N)
	A := make([]int, N)
	B := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(_r, &A[i], &B[i])
	}
	ans := Solve(N, A, B)
	fmt.Fprintf(_w, "%d\n", ans)
}

func Solve(N int, A, B []int) int {
	C := make([][]int, N)
	for i := 0; i < N; i++ {
		C[i] = []int{A[i], B[i]}
	}
	sort.Slice(C, func(i, j int) bool { return C[i][0]+C[i][1] > C[j][0]+C[j][1] })
	ret := 0
	for i := 0; i < N; i++ {
		if i%2 == 0 {
			ret += C[i][0]
		} else {
			ret -= C[i][1]
		}
	}
	return ret
}

var _r, _w = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
