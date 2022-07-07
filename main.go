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
	// MOD = int(1e9 + 7)
	MOD = 998244353
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
	N, M := ni(), ni()
	edges := make([][]Edge, N)
	for i := 0; i < M; i++ {
		A, B, C := ni()-1, ni()-1, ni()
		edges[A] = append(edges[A], Edge{B, C})
		edges[B] = append(edges[B], Edge{A, C})
	}
	ans := Solve(N, M, edges)
	fmt.Fprintf(wr, "%d\n", ans)
}

func Solve(N, M int, edges [][]Edge) int {
	ret := 0
	for i := 0; i < N; i++ {
		for _, e := range edges[i] {
			ret += MaxInt(0, e.cost)
		}
	}
	ret /= 2
	pq := make(PQueue, 0)
	heap.Init(&pq)
	used := make([]bool, N)
	used[0] = true
	for i := 0; i < len(edges[0]); i++ {
		e := edges[0][i]
		if used[e.to] {
			continue
		}
		heap.Push(&pq, &e)
	}
	for pq.Len() > 0 {
		e := heap.Pop(&pq).(*Edge)
		if used[e.to] {
			continue
		} else {
			ret -= MaxInt(0, e.cost)
			used[e.to] = true
			for i := 0; i < len(edges[e.to]); i++ {
				f := edges[e.to][i]
				if used[f.to] {
					continue
				}
				heap.Push(&pq, &f)
			}
		}
	}
	return ret
}

type PQueue []*Edge

func (pq PQueue) Len() int           { return len(pq) }
func (pq PQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }
func (pq PQueue) Less(i, j int) bool { return pq[i].cost < pq[j].cost }

func (pq *PQueue) Push(x interface{}) {
	item := x.(*Edge)
	*pq = append(*pq, item)
}
func (pq *PQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	*pq = old[:n-1]
	item := old[n-1]
	old[n-1] = nil
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

func ni() int {
	sc.Scan()
	x, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return x
}

func nl() string {
	sc.Scan()
	return sc.Text()
}
