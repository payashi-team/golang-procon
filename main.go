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

type Query struct {
	c, x int
}

func main() {
	defer _w.Flush()
	var Q int
	fmt.Fscan(_r, &Q)
	queries := make([]Query, Q)
	for i := 0; i < Q; i++ {
		var c, x int
		fmt.Fscan(_r, &c)
		if c == 1 {
			fmt.Fscan(_r, &x)
		}
		queries[i] = Query{c, x}

	}
	Solve(Q, queries)
}

func Solve(Q int, queries []Query) {
	que := make([]int, 0)
	pq := make(PQueue, 0)
	heap.Init(&pq)
	for _, q := range queries {
		switch q.c {
		case 1:
			que = append(que, q.x)
		case 2:
			if pq.Len() > 0 {
				item := heap.Pop(&pq).(*Item)
				fmt.Fprintf(_w, "%d\n", item.val)
			} else {
				fmt.Fprintf(_w, "%d\n", que[0])
				que = que[1:]
			}
		case 3:
			for _, v := range que {
				heap.Push(&pq, &Item{-v, v, -1})
			}
			que = []int{}
		}
	}
}

type Item struct {
	priority, val, index int
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
	item.index = -1
	old[n-1] = nil
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
