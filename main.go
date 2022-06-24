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
	to, yen, snook int
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	N, M, S, T := NextInt(), NextInt(), NextInt(), NextInt()
	S--
	T--
	edges := make([][]Edge, N)
	for i := 0; i < M; i++ {
		u, v, a, b := NextInt(), NextInt(), NextInt(), NextInt()
		u--
		v--
		edges[u] = append(edges[u], Edge{v, a, b})
		edges[v] = append(edges[v], Edge{u, a, b})
	}
	ans := Solve(N, M, S, T, edges)
	for _, v := range ans {
		fmt.Fprintf(wr, "%d\n", v)
	}
}

func Solve(N, M, S, T int, edges [][]Edge) []int {
	dist1 := make([]int, N) // from s by yen
	for i := 0; i < N; i++ {
		dist1[i] = INF
	}
	dist1[S] = 0
	pq := make(PQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &Item{S, 0})
	used := make([]bool, N)
	// dijkstra 1
	for pq.Len() > 0 {
		p := heap.Pop(&pq).(*Item)
		used[p.node] = true
		if dist1[p.node] < p.cost {
			continue
		}
		for _, e := range edges[p.node] {
			if !used[e.to] && dist1[e.to] > p.cost+e.yen {
				dist1[e.to] = p.cost + e.yen
				heap.Push(&pq, &Item{e.to, p.cost + e.yen})
			}
		}
	}
	dist2 := make([]int, N) // from t by snook
	for i := 0; i < N; i++ {
		dist2[i] = INF
	}
	dist2[T] = 0
	heap.Push(&pq, &Item{T, 0})
	used = make([]bool, N)
	// dijkstra 2
	for pq.Len() > 0 {
		p := heap.Pop(&pq).(*Item)
		used[p.node] = true
		if dist2[p.node] < p.cost {
			continue
		}
		for _, e := range edges[p.node] {
			if !used[e.to] && dist2[e.to] > p.cost+e.snook {
				dist2[e.to] = p.cost + e.snook
				heap.Push(&pq, &Item{e.to, p.cost + e.snook})
			}
		}
	}
	ret := make([]int, N)
	for i := 0; i < N; i++ {
		ret[i] = int(1e15) - dist1[i] - dist2[i]
	}
	for i := N - 2; i >= 0; i-- {
		ret[i] = MaxInt(ret[i], ret[i+1])
	}
	return ret

}

type Item struct {
	node, cost int
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
