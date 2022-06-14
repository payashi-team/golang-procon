package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

const (
	INF = int(1 << 61)
	// MOD = int(1e9 + 7)
	MOD = 998244353
)

var _w, _r, _s = bufio.NewWriter(os.Stdout), bufio.NewReader(os.Stdin), bufio.NewScanner(os.Stdout)

func main() {
	defer _w.Flush()
	var N, M, Q int
	fmt.Fscan(_r, &N, &M, &Q)
	items := make([]Item, N)
	for i := 0; i < N; i++ {
		item := &items[i]
		fmt.Fscan(_r, &item.w, &item.v)
	}
	sort.Slice(items, func(i, j int) bool {
		if items[i].v != items[j].v {
			return items[i].v > items[j].v
		} else {
			return items[i].w > items[j].w
		}
	})
	X := make([]Box, M)
	for i := 0; i < M; i++ {
		fmt.Fscan(_r, &X[i].cap)
		X[i].idx = i
	}
	sort.Slice(X, func(i, j int) bool {
		return X[i].cap < X[j].cap
	})
	for q := 0; q < Q; q++ {
		var l, r int
		fmt.Fscan(_r, &l, &r)
		l--
		r--
		ret := 0
		used := make([]bool, N)
		for i := 0; i < M; i++ {
			if l <= X[i].idx && X[i].idx <= r {
				continue
			}
			for j := 0; j < N; j++ {
				if used[j] {
					continue
				}
				if X[i].cap >= items[j].w {
					used[j] = true
					ret += items[j].v
					break
				}
			}
		}
		// fmt.Fprintf(_w, "%d\n", ret)
		fmt.Printf("%d\n", ret)
	}
}

type Item struct {
	w, v int
}

type Box struct {
	cap, idx int
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

func nextInt() int {
	_s.Scan()
	i, e := strconv.Atoi(_s.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func nextLine() string {
	_s.Scan()
	return _s.Text()
}
