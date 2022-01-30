package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const (
	INF = int(1 << 62)
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
		fmt.Fscan(_r, &A[i], &B[i])
	}
	ans := Solve(N, A, B)
	for i := 1; i <= N; i++ {
		fmt.Fprintf(_w, "%d ", ans[i])
	}
	fmt.Fprintln(_w)
}

type Item struct {
	date, value int
}

func Solve(N int, A, B []int) map[int]int {
	items := make([]Item, N*2)
	for i := 0; i < N; i++ {
		items[i*2] = Item{A[i], 1}
		items[i*2+1] = Item{A[i] + B[i], -1}
	}
	sort.Slice(items, func(i, j int) bool { return items[i].date < items[j].date })
	ret := make(map[int]int)
	cur := Item{1, 0}
	for _, item := range items {
		ret[cur.value] += item.date - cur.date
		cur.value += item.value
		cur.date = item.date
	}
	return ret
}

var _r, _w = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
