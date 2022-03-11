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
	command string
	n, m    int
}

func main() {
	defer _w.Flush()
	var Q, L int
	fmt.Fscan(_r, &Q, &L)
	Solve(Q, L)
}

type Item struct {
	val, acc int
}

func Solve(Q, L int) {
	size := 0
	stack := make([]Item, 0)
	safe := true
	for i := 0; i < Q; i++ {
		var s string
		var n, m int
		fmt.Fscan(_r, &s)
		if s == "Push" {
			fmt.Fscan(_r, &n, &m)
		} else if s == "Pop" {
			fmt.Fscan(_r, &n)
		}
		q := Query{s, n, m}
		if !safe {
			continue
		}
		if q.command == "Push" {
			size += q.n
			if size > L {
				fmt.Fprintf(_w, "FULL\n")
				safe = false
				continue
			}
			stack = append(stack, Item{q.m, size})
		} else if q.command == "Pop" {
			size -= q.n
			if size < 0 {
				fmt.Fprintf(_w, "EMPTY\n")
				safe = false
				continue
			}
			idx := sort.Search(len(stack), func(j int) bool { return stack[j].acc >= size })
			stack = stack[:idx+1]
			stack[idx].acc = size
		} else if q.command == "Top" {
			if size == 0 {
				fmt.Fprintf(_w, "EMPTY\n")
				safe = false
				continue
			}
			fmt.Fprintf(_w, "%d\n", stack[len(stack)-1].val)
		} else if q.command == "Size" {
			fmt.Fprintf(_w, "%d\n", size)
		}
		// fmt.Fprintf(_w, "%v\n", stack)
	}
	if safe {
		fmt.Fprintf(_w, "SAFE\n")
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
