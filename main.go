package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
)

const (
	INF = int(1 << 61)
	MOD = int(1e9 + 7)
	// MOD = 998244353
)

type Edge struct {
	to, cost int
}

type Item struct {
	to, cost, priority, index int
}

type PQueue []*Item

func main() {
	defer _w.Flush()
	var N, M, S, T int
	fmt.Fscan(_r, &N, &M, &S, &T)
	D := make([][]Edge, N)
	for i := 0; i < M; i++ {
		var x, y, d int
		fmt.Fscan(_r, &x, &y, &d)
		x--
		y--
		D[x] = append(D[x], Edge{y, d})
		D[y] = append(D[y], Edge{x, d})
	}
	ans := Solve(N, M, S, T, D)
	fmt.Fprintf(_w, "%d\n", ans)
}

func Solve(N, M, S, T int, D [][]Edge) int {
	S--
	T--
	dijkstra := func(s int) []int {
		ds := make([]int, N)
		for i := 0; i < N; i++ {
			ds[i] = INF
		}
		ds[s] = 0
		pq := make(PQueue, 0)
		heap.Init(&pq)
		heap.Push(&pq, &Item{s, 0, -0, -1})
		for pq.Len() > 0 {
			item := pq.Pop().(*Item)
			if ds[item.to] < item.cost {
				continue
			}
			for _, e := range D[item.to] {
				if ds[e.to] > ds[item.to]+e.cost {
					ds[e.to] = ds[item.to] + e.cost
					heap.Push(&pq, &Item{e.to, ds[e.to], -ds[e.to], -1})
				}
			}
		}
		return ds
	}
	d1 := dijkstra(S)
	d2 := dijkstra(T)
	for i := 0; i < N; i++ {
		if d1[i]!=INF &&d1[i] == d2[i] {
			return i + 1
		}
	}
	return -1
}

func (pq PQueue) Len() int { return len(pq) }
func (pq PQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}
func (pq PQueue) Less(i, j int) bool {
	return pq[i].priority > pq[j].priority
}

func (pq *PQueue) Push(x interface{}) {
	item := x.(*Item)
	item.index = len(*pq)
	*pq = append(*pq, item)
}

func (pq *PQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[:n-1]
	return item
}

func AbsInt(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

func MaxInt(nums ...int) int {
	ret := -1
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
