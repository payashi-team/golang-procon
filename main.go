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
	var N, M int
	fmt.Fscanf(_r, "%d %d\n", &N, &M)
	a, b, c := Solve(N, M)
	fmt.Fprintf(_w, "%d %d %d\n", a, b, c)
}

func Solve(N, M int) (int, int, int) {
	if M < 2*N || 4*N < M {
		return -1, -1, -1
	}
	for a := 0; a <= M/2; a++ {
		b := -M + 4*N - 2*a
		c := M - 3*N + a
		if 0 <= b && 0 <= c {
			return a, b, c
		}
	}
	return -1, -1, -1
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
