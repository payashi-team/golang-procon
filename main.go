package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

const (
	INF = int(1 << 61)
	// MOD = int(1e9 + 7)
	MOD = 998244353
)

func main() {
	defer _w.Flush()
	var N, M int
	fmt.Fscan(_r, &N, &M)
	A := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(_r, &A[i])
	}
	ans := Solve(N, M, A)
	fmt.Fprintf(_w, "%d\n", ans)
}

func Solve(N, M int, A []int) int {
	pos:=make([][]int, int(15e5)+1)
	for i := 0; i <= int(15e5); i++ {
		pos[i] = make([]int, 0)
	}
	for i, v := range A {
		pos[v] = append(pos[v], i+1)
	}
	for i := 0; i <= int(15e5); i++ {
		if len(pos[i])==0{
			return i
		}
		pos[i] = append([]int{0}, pos[i]...)
		pos[i] = append(pos[i], N+1)
		// fmt.Printf("%d: %v\n", i, pos[i])
		for j := 0; j < len(pos[i])-1; j++ {
			if pos[i][j+1]-pos[i][j]>M{
				return i
			}
		}
	}
	return int(15e5)+1
}

func Contains(x int, nums ...int) bool {
	for _, v := range nums {
		if v == x {
			return true
		}
	}
	return false
}

func AbsInt(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

func MaxInt(nums ...int) int {
	ret := -INF
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
