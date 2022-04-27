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
	ans := Solve(N)
	fmt.Fprintf(_w, "%.10f\n", ans)
}

func Solve(N int) float64 {
	n := float64(N)
	var ret float64
	for i := N - 1; i >= 1; i-- {
		ret += n / float64(i)
	}
	return ret
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
