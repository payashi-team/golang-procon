package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	defer _w.Flush()
	var N int
	fmt.Fscan(_r, &N)
	A := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(_r, &A[i])
	}
	ans := Solve(N, A)
	for _, v := range ans {
		fmt.Fprintf(_w, "%d\n", v)
	}
}

func Solve(N int, A []int) []int {
	ret:=make([]int, 0)
	B := make([]int, N)
	for cnt:=0;cnt<N;cnt++{
		update:=false
		for i := N-1; i >= 0; i--{
			if A[i]-B[i]==1{
				update = true
				ret = append(ret, A[i])
				A[i]=-1
				for j := i+1; j < N; j++ {
					B[j]++
				}
				break
			}
		}
		if !update{
			return []int{-1}
		}
	}
	return ret
}

var _r, _w = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
