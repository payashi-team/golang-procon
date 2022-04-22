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
	// MOD = int(1e9 + 7)
	MOD = 998244353
)

type Edge struct {
	to, cost, k int
}

func main() {
	defer _w.Flush()
	var N, M, X, Y int
	fmt.Fscan(_r, &N, &M, &X, &Y)
	X--
	Y--
	edges := make([][]Edge, N)
	for i := 0; i < N; i++ {
		edges[i] = make([]Edge, 0)
	}
	for i := 0; i < M; i++ {
		var a, b, t, k int
		fmt.Fscan(_r, &a, &b, &t, &k)
		a--
		b--
		edges[a] = append(edges[a], Edge{b, t, k})
		edges[b] = append(edges[b], Edge{a, t, k})
	}
	ans := Solve(N, M, X, Y, edges)
	fmt.Fprintf(_w, "%d\n", ans)
}

func Solve(N, M, X, Y int, edges [][]Edge) int {
	dist := make([]int, N)
	for i := 0; i < N; i++ {
		dist[i] = INF
	}
	dist[X] = 0
	pq := make(PQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &Item{X, 0, -1})
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		if item.dist > dist[item.to] {
			continue
		}
		for _, e := range edges[item.to] {
			leave := item.dist + (e.k-(item.dist%e.k))%e.k
			if dist[e.to] > leave+e.cost {
				dist[e.to] = leave + e.cost
				heap.Push(&pq, &Item{e.to, dist[e.to], -1})
			}
		}
	}
	if dist[Y] == INF {
		return -1
	} else {
		return dist[Y]
	}
}

type Item struct {
	to, dist, index int
}

type PQueue []*Item

func (pq PQueue) Len() int { return len(pq) }
func (pq PQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}
func (pq PQueue) Less(i, j int) bool { return pq[i].dist < pq[j].dist }

func (pq *PQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1
	(*pq)[n-1] = nil
	*pq = old[:n-1]
	return item
}

func (pq *PQueue) Push(x interface{}) {
	item := x.(*Item)
	n := len(*pq)
	item.index = n
	*pq = append(*pq, item)
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
