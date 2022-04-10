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
	MOD = int(1e9 + 7)
	// MOD = 998244353
)

type Query struct {
	mode, x, c int
}

func main() {
	defer _w.Flush()
	var Q int
	fmt.Fscan(_r, &Q)
	Qs := make([]Query, Q)
	for i := 0; i < Q; i++ {
		var mode, x, c int
		fmt.Fscan(_r, &mode)
		if mode == 1 {
			fmt.Fscan(_r, &x, &c)
			Qs[i] = Query{mode, x, c}
		} else {
			fmt.Fscan(_r, &c)
			Qs[i] = Query{mode, x, c}
		}
	}
	Solve(Q, Qs)
}

type Item struct {
	x, l, r int
}

type Flag struct {
	block, num int
}

func Solve(Q int, Qs []Query) {
	sum := 0
	outsum := 0
	lb := Flag{0, 0}
	que := make([]Item, 0)
	for _, q := range Qs {
		// fmt.Printf("%v\n", que)
		if q.mode == 1 {
			que = append(que, Item{q.x, sum, sum + q.c})
			sum += q.c
		} else {
			outsum += q.c
			idx := sort.Search(len(que), func(i int) bool { return que[i].r >= outsum })
			ub := Flag{idx, outsum - que[idx].l}
			// fmt.Printf("lb: %d block, %d\nub: %d block, %d\n", lb.block, lb.num, ub.block, ub.num)
			// fmt.Printf("outsum: %d\n", outsum)
			ret := 0
			if lb.block == ub.block {
				ret += que[idx].x * (ub.num - lb.num)
			} else {
				ret += que[lb.block].x * (que[lb.block].r - que[lb.block].l - lb.num)
				ret += que[ub.block].x * ub.num
				for i := lb.block + 1; i < ub.block; i++ {
					ret += que[i].x * (que[i].r - que[i].l)
				}
			}
			lb = ub
			fmt.Printf("%d\n", ret)
		}
	}
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
