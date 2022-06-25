package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
	"strconv"
)

const (
	INF = int(1 << 61)
	MOD = int(1e9 + 7)
	// MOD = 998244353
)

var sc = bufio.NewScanner(os.Stdin)
var wr = bufio.NewWriter(os.Stdout)

type Edge struct {
	to, cost int
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	N, M, K := NextInt(), NextInt(), NextInt()
	edges := make([][]Edge, N)
	X := make([]int, N)
	Y := make([]int, N)
	for i := 0; i < M; i++ {
		a, b, c := NextInt(), NextInt(), NextInt()
		a--
		b--
		edges[a] = append(edges[a], Edge{b, c})
		edges[b] = append(edges[b], Edge{a, c})
	}
	for i := 0; i < N; i++ {
		X[i], Y[i] = NextInt(), NextInt()
	}
	ans := Solve(N, M, K, edges, X, Y)
	fmt.Fprintf(wr, "%d\n", ans)
}

func Solve(N, M, K int, edges [][]Edge, X, Y []int) int {
	pq := make(PQueue, 0)
	heap.Init(&pq)
	used := make([]bool, N*2*K)
	dist := make([]int, N*2*K)
	for i := 0; i < N*2*K; i++ {
		dist[i] = INF
	}
	dist[0] = 0
	heap.Push(&pq, &Item{0, 0, 0})
	for pq.Len() > 0 {
		p := heap.Pop(&pq).(*Item)
		u := p.flower*N + p.node
		if dist[u] < p.cost {
			continue
		}
		used[u] = true
		for _, e := range edges[p.node] {
			v := p.flower*N + e.to
			if !used[v] && dist[v] > dist[u]+e.cost {
				dist[v] = dist[u] + e.cost
				heap.Push(&pq, &Item{e.to, dist[v], p.flower})
			}
		}
		if p.flower+X[p.node]>=2*K{
			continue
		}
		t := (p.flower+X[p.node])*N + p.node
		if !used[t] && p.flower < K && dist[t] > dist[u]+Y[p.node] {
			dist[t] = dist[u] + Y[p.node]
			heap.Push(&pq, &Item{p.node, dist[t], p.flower + X[p.node]})
		}
	}
	ret := INF
	for i := 0; i < K; i++ {
		ret = MinInt(ret, dist[(K+i)*N+N-1])
	}
	if ret == INF {
		return -1
	}
	return ret

}

type Item struct {
	node, cost, flower int
}

type PQueue []*Item

func (pq PQueue) Len() int { return len(pq) }
func (pq PQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq PQueue) Less(i, j int) bool {
	return pq[i].cost < pq[j].cost
}
func (pq *PQueue) Pop() interface{} {
	old := (*pq)
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	*pq = old[:n-1]
	return item
}
func (pq *PQueue) Push(x interface{}) {
	item := x.(*Item)
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

func NextInt() int {
	sc.Scan()
	x, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return x
}

func NextLine() string {
	sc.Scan()
	return sc.Text()
}
