package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
	"sort"
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
	A := make([]int, M)
	B := make([]int, M)
	for i := 0; i < M; i++ {
		fmt.Fscan(_r, &A[i], &B[i])
	}
	ans, ok := Solve(N, M, A, B)
	if !ok {
		fmt.Fprintf(_w, "-1\n")
	} else {
		for _, v := range ans {
			fmt.Fprintf(_w, "%d ", v)
		}
		fmt.Fprintln(_w)
	}
}

func Solve(N, M int, A, B []int) ([]int, bool) {
	ret := make([]int, 0)
	pq := make(PQueue, 0)
	heap.Init(&pq)
	ins := make([]int, N) // 入次数
	edges := make([][]int, N)
	for i := 0; i < M; i++ {
		a := A[i] - 1
		b := B[i] - 1
		ins[b]++
		edges[a] = append(edges[a], b)
	}
	for i := 0; i < N; i++ {
		sort.Ints(edges[i])
	}
	for i := 0; i < N; i++ {
		if ins[i] == 0 {
			heap.Push(&pq, &Item{i})
		}
	}
	for pq.Len() > 0 {
		u := heap.Pop(&pq).(*Item).pos
		ret = append(ret, u+1)
		for _, v := range edges[u] {
			ins[v]--
			if ins[v] == 0 {
				heap.Push(&pq, &Item{v})
			}
		}
	}
	// 閉路がある
	if len(ret) < N {
		return ret, false
	} else {
		return ret, true
	}
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
