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
	fmt.Fscanf(_r, "%d\n", &N)
	R := make([]int, 0)
	G := make([]int, 0)
	B := make([]int, 0)
	for i := 0; i < 2*N; i++ {
		var v int
		var c rune
		fmt.Fscanf(_r, "%d %c\n", &v, &c)
		if c == 'R' {
			R = append(R, v)
		} else if c == 'G' {
			G = append(G, v)
		} else {
			B = append(B, v)
		}
	}
	ans := Solve(N, R, G, B)
	fmt.Fprintf(_w, "%d\n", ans)
}

func Solve(N int, R, G, B []int) int {
	sort.Ints(R)
	sort.Ints(G)
	sort.Ints(B)
	if len(R)%2+len(G)%2+len(B)%2 == 0 {
		return 0
	}
	// make len(R) even
	if len(G)%2 == 0 {
		R, G = G, R
	} else if len(B)%2 == 0 {
		R, B = B, R
	}
	rg := MinDist(R, G)
	rb := MinDist(R, B)
	gb := MinDist(G, B)
	return MinInt(rg+rb, gb)
}

// Both a and b are sorted
func MinDist(a, b []int) int {
	ret := INF
	for _, v := range a {
		idx := sort.Search(len(b), func(i int) bool { return v <= b[i] })
		if idx == 0 {
			ret = MinInt(ret, b[idx]-v)
		} else if idx == len(b) {
			ret = MinInt(ret, v-b[idx-1])
		} else {
			ret = MinInt(ret, b[idx]-v, v-b[idx-1])
		}
	}
	return ret
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
