package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const (
	INF = int(1 << 60)
	// MOD = int(1e9 + 7)
	MOD = 998244353
)

func main() {
	defer _w.Flush()
	var T int
	fmt.Fscan(_r, &T)
	R := make([]int, T)
	G := make([]int, T)
	B := make([]int, T)
	for i := 0; i < T; i++ {
		fmt.Fscan(_r, &R[i], &G[i], &B[i])
	}
	ans := Solve(T, R, G, B)
	for _, v := range ans {
		fmt.Fprintf(_w, "%d\n", v)
	}
}

func Solve(T int, R, G, B []int) []int {
	ret := make([]int, T)
	for i := 0; i < T; i++ {
		ret[i] = INF
		arr := []int{R[i], G[i], B[i]}
		sort.Ints(arr)
		if (arr[2]-arr[1])%3 == 0 {
			ret[i] = MinInt(ret[i], arr[2])
		}
		if (arr[2]-arr[0])%3 == 0 {
			ret[i] = MinInt(ret[i], arr[2])
		}
		if (arr[1]-arr[0])%3 == 0 {
			ret[i] = MinInt(ret[i], arr[1])
		}
		if ret[i] == INF {
			ret[i] = -1
		}
	}
	return ret
}

func MinInt(nums ...int) int {
	ret := INF
	for _, v := range nums {
		if ret > v {
			ret = v
		}
	}
	return ret
}

var _r, _w = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
