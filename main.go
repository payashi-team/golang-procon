package main

import (
	"bufio"
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
	var N int
	fmt.Fscan(_r, &N)
	A := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(_r, &A[i])
	}
	ans := Solve(N, A)
	fmt.Fprintf(_w, "%d\n", ans)
}

func Solve(N int, A []int) int {
	// Ai + Aj = j-i (i<j)
	// Ai+i + Aj-j = 0 (i<j)
	// Ai-i + Aj+j = 0
	// Bi = Ai+i
	// Ci = Ai-i
	// Bi+Cj=0 (i!=j)
	// O(NlogN)
	B := make([]int, N)
	C := make([]int, N)
	for i := 0; i < N; i++ {
		B[i] = A[i] + (i + 1)
		C[i] = A[i] - (i + 1)
	}
	sort.Ints(B)
	sort.Ints(C)
	ans := 0
	for _, b := range B {
		// cnt := sort.SearchInts(C, -b+1) - sort.SearchInts(C, -b)
		cnt :=
			sort.Search(N, func(i int) bool {
				return C[i] > -b
			}) -
				sort.Search(N, func(i int) bool {
					return C[i] >= -b
				})
		// fmt.Printf("%d in C = %d\n", -b, cnt)
		ans += cnt
	}
	return ans
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
