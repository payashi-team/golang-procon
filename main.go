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
		fmt.Fscan(_r, &A[i])
	}
	for i := 0; i < N; i++ {
		fmt.Fscan(_r, &B[i])
	}
	ans := Solve(N, A, B)
	fmt.Fprintf(_w, "%d\n", ans)
}

func Solve(N int, A, B []int) int {
	sumA := 0
	sumB := 0
	for i := 0; i < N; i++ {
		sumA += A[i]
		sumB += B[i]
	}
	if sumA < sumB {
		return -1
	}
	ret := 0
	shortage := 0
	surplus := make([]int, 0)
	for i := 0; i < N; i++ {
		if A[i] < B[i] {
			shortage += B[i] - A[i]
			ret++
		} else if A[i] > B[i] {
			surplus = append(surplus, A[i]-B[i])
		}
	}
	sort.Slice(surplus, func(i, j int) bool {return surplus[i]>surplus[j]})
	for i:=0;shortage>0;i++ {
		shortage -=surplus[i]
		ret++
	}
	return ret
}

var _r, _w = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
