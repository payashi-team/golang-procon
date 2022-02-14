package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

const (
	INF = int(1 << 60)
	// MOD = int(1e9 + 7)
	MOD = 998244353
)

func main() {
	defer _w.Flush()
	var N int
	fmt.Fscan(_r, &N)
	edges := make([][]int, N)
	for i := 0; i < N-1; i++ {
		var a, b int
		fmt.Fscan(_r, &a, &b)
		a--
		b--
		edges[a] = append(edges[a], b)
		edges[b] = append(edges[b], a)
	}
	ans := Solve(N, edges)
	for i, num := range ans {
		fmt.Fprintf(_w, "%d", num)
		if i == N-1 {
			fmt.Fprintln(_w)
		} else {
			fmt.Fprintf(_w, " ")
		}
	}
}

type Item struct {
	value, priority, index int
}

type PQueue []*Item

func (pq PQueue) Len() int {
	return len(pq)
}

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
	*pq = old[0 : n-1]
	return item
}

func Solve(N int, edges [][]int) []int {
	pq := make(PQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &Item{0, 0, -1})
	ret := make([]int, 0)
	used := make([]bool, N)
	for pq.Len() > 0 {
		p := heap.Pop(&pq).(*Item)
		used[p.value] = true
		ret = append(ret, p.value)
		for _, v := range edges[p.value] {
			if used[v] {
				continue
			}
			heap.Push(&pq, &Item{v, -v, -1})
		}
	}
	for i := 0; i < N; i++ {
		ret[i]++
	}
	return ret
}

var _r, _w = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
