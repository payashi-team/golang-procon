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
	var N, K int
	S := make([]byte, N)
	fmt.Fscan(_r, &N, &K, &S)
	ans := Solve(N, K, S)
	fmt.Fprintf(_w, "%c\n", ans)
}

func Solve(N, K int, S []byte) byte {
	win := func(a, b byte) byte {
		if a == b {
			return a
		}
		var r byte = 'R'
		var p byte = 'P'
		var s byte = 'S'

		if a+b == r+p {
			return p
		} else if a+b == r+s {
			return r
		} else {
			return s
		}
	}
	for i := 0; i < K; i++ {
		T := make([]byte, 2*N)
		for j := 0; j < N; j++ {
			T[j] = S[j]
			T[j+N] = S[j]
		}
		for j := 0; j < N; j++ {
			S[j] = win(T[2*j+1], T[2*j])
		}
	}
	return S[0]
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
