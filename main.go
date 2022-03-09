package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

const (
	INF = int(1 << 61)
	MOD = int(1e9 + 7)
	// MOD = 998244353
)

type Point struct {
	h, w int
}

func main() {
	defer _w.Flush()
	var H, W, M int
	fmt.Fscan(_r, &H, &W, &M)
	targets := make([]Point, M)
	for i := 0; i < M; i++ {
		fmt.Fscan(_r, &targets[i].h, &targets[i].w)
	}
	ans := Solve(H, W, M, targets)
	fmt.Fprintf(_w, "%d\n", ans)
}

func Solve(H, W, M int, targets []Point) int {
	cnth := make([]int, H)
	cntw := make([]int, W)
	filled := make(map[Point]bool)
	for _, t := range targets {
		cnth[t.h-1]++
		cntw[t.w-1]++
		filled[t] = true
	}
	maxh := -1
	hs := make([]int, 0)
	for i := 0; i < H; i++ {
		if cnth[i] >= maxh {
			if cnth[i] > maxh {
				hs = []int{}
			}
			maxh = cnth[i]
			hs = append(hs, i)
		}
	}
	maxw := -1
	ws := make([]int, 0)
	for i := 0; i < W; i++ {
		if cntw[i] >= maxw {
			if cntw[i] > maxw {
				ws = []int{}
			}
			maxw = cntw[i]
			ws = append(ws, i)
		}
	}
	ret := maxh + maxw - 1
	for _, h := range hs {
		for _, w := range ws {
			if !filled[Point{h + 1, w + 1}] {
				ret++
				return ret
			}
		}
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
