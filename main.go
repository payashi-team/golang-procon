package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

const (
	INF = int(1 << 60)

// MOD = int(1e9 + 7)
// MOD = 998244353
)

func main() {
	defer _w.Flush()
	b := make(map[int]int)
	for i := 0; i < 10; i++ {
		var tmp int
		fmt.Scan(&tmp)
		b[tmp] = i
	}
	var N int
	fmt.Fscan(_r, &N)
	A := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(_r, &A[i])
	}
	ans := Solve(b, N, A)
	for _, v := range ans {
		fmt.Fprintf(_w, "%d\n", v)
	}
}

type Item struct {
	normal, abnormal int
}

func Solve(b map[int]int, N int, A []int) []int {
	B := make([]Item, N)
	for i := 0; i < N; i++ {
		abnormal := strconv.Itoa(A[i])
		normal := 0
		for _, c := range abnormal {
			v, _ := strconv.Atoi(string(c))
			normal *= 10
			normal += b[v]
		}
		B[i] = Item{normal, A[i]}
	}
	sort.Slice(B, func(i, j int) bool { return B[i].normal < B[j].normal })
	ret := make([]int, N)
	for i := 0; i < N; i++ {
		ret[i] = B[i].abnormal
	}
	return ret
}

var _r, _w = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
