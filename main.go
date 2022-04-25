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
	var N int
	fmt.Fscan(_r, &N)
	edges := make([][]int, N)
	for i := 0; i < N-1; i++ {
		var u, v int
		fmt.Fscan(_r, &u, &v)
		u--
		v--
		edges[u] = append(edges[u], v)
		edges[v] = append(edges[v], u)
	}
	ans := Solve(N, edges)
	for _, v := range ans {
		fmt.Fprintf(_w, "%d %d\n", v[0], v[1])
	}
}

func Solve(N int, edges [][]int) [][]int {
	used := make([]bool, N)
	lr := make([][]int, N)
	cnt := 1
	var dfs func(int)
	dfs = func(u int) {
		used[u] = true
		l := N + 1
		r := -1
		for _, v := range edges[u] {
			if used[v] {
				continue
			}
			dfs(v)
			l = MinInt(l, lr[v][0])
			r = MaxInt(r, lr[v][1])
		}
		if r < 0 {
			lr[u] = []int{cnt, cnt}
			cnt++
			return
		}
		lr[u] = []int{l, r}
	}
	dfs(0)
	return lr
}

type Item struct {
	pos int
}

type PQueue []*Item

func (pq PQueue) Len() int           { return len(pq) }
func (pq PQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }
func (pq PQueue) Less(i, j int) bool { return pq[i].pos < pq[j].pos }

func (pq *PQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*Item))
}

func (pq *PQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[:n-1]
	return item
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
