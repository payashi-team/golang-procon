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

func main() {
	defer _w.Flush()
	var N, M int
	fmt.Fscan(_r, &N, &M)
	H := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(_r, &H[i])
	}
	U := make([]int, M)
	V := make([]int, M)
	for i := 0; i < M; i++ {
		fmt.Fscan(_r, &U[i], &V[i])
	}
	ans := Solve(N, M, H, U, V)
	fmt.Fprintf(_w, "%d\n", ans)
}

type Edge struct {
	to, cost int
}

func Solve(N, M int, H, U, V []int) int {
	// tanoshisa + hyoukou
	cost := make([][]Edge, N)

	for i := 0; i < M; i++ {
		u := U[i] - 1
		v := V[i] - 1
		diff := H[u] - H[v]
		if diff > 0 {
			cost[u] = append(cost[u], Edge{v, 0})
			cost[v] = append(cost[v], Edge{u, diff})
		} else {
			cost[u] = append(cost[u], Edge{v, -diff})
			cost[v] = append(cost[v], Edge{u, 0})
		}
	}

	dist := make([]int, N)
	for i := 0; i < N; i++ {
		dist[i] = INF
	}
	dist[0] = 0
	pq := make(PQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &Item{0, 0, -1})
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		u := item.value
		if dist[u] < -item.priority {
			continue
		}
		for _, e := range cost[u] {
			if dist[e.to] > dist[u]+e.cost {
				dist[e.to] = dist[u] + e.cost
				heap.Push(&pq, &Item{e.to, -dist[e.to], -1})
			}
		}
	}
	ans := 0
	for i := 0; i < N; i++ {
		ans = MaxInt(ans, H[0]-H[i]-dist[i])
	}
	return ans
}

type Item struct {
	value, priority, index int
}

type PQueue []*Item

func (pq PQueue) Len() int {
	return len(pq)
}

func (pq PQueue) Less(i, j int) bool {
	return pq[i].priority > pq[j].priority
}

func (pq PQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := (*pq)[n-1]
	item.index = -1
	(*pq)[n-1] = nil
	*pq = (*pq)[:n-1]
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
