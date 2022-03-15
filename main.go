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

func main() {
	defer _w.Flush()
	var N int
	fmt.Fscan(_r, &N)
	R := make([]int, N)
	H := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(_r, &R[i], &H[i])
	}
	ans := Solve(N, R, H)
	for _, res := range ans {
		fmt.Fprintf(_w, "%d %d %d\n", res.win, res.lose, res.draw)
	}
}

type Result struct {
	win, lose, draw int
}

type Player struct {
	rate, hand, index int
}

func Solve(N int, R, H []int) []Result {
	ret := make([]Result, N)
	ps := make([]Player, N)
	for i := 0; i < N; i++ {
		ps[i] = Player{R[i], H[i] - 1, i}
	}
	sort.Slice(ps, func(i, j int) bool { return ps[i].rate < ps[j].rate })
	S := make([][3]int, N)
	S[0][ps[0].hand]++
	for i := 1; i < N; i++ {
		S[i] = S[i-1]
		S[i][ps[i].hand]++
	}
	for i := 0; i < N; i++ {
		p := ps[i]
		ub := sort.Search(N, func(j int) bool { return ps[j].rate > p.rate })
		lb := sort.Search(N, func(j int) bool { return ps[j].rate >= p.rate })
		r := &ret[p.index]
		r.win = lb
		r.lose = N - ub
		for j := 0; j < 3; j++ {
			var cnt int
			if lb == 0 {
				cnt = S[ub-1][j]
			} else {
				cnt = S[ub-1][j] - S[lb-1][j]
			}
			switch (p.hand - j + 3) % 3 {
			case 0:
				r.draw += cnt
			case 1:
				r.lose += cnt
			case 2:
				r.win += cnt
			}
		}
		r.draw--
	}
	return ret
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
