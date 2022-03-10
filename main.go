package main

import (
	"bufio"
	"fmt"
	"math"
	"math/bits"
	"os"
)

const (
	INF = int(1 << 61)
	MOD = int(1e9 + 7)
	// MOD = 998244353
)

func main() {
	defer _w.Flush()
	var S string
	var Q int
	fmt.Fscan(_r, &S, &Q)
	T := make([]int, Q)
	K := make([]int, Q)
	for i := 0; i < Q; i++ {
		fmt.Fscan(_r, &T[i], &K[i])
	}
	Solve(S, Q, T, K)
}

func Solve(S string, Q int, T, K []int) {
	mp := map[byte]int{'A': 0, 'B': 1, 'C': 2}
	for i := 0; i < Q; i++ {
		t := T[i]
		k := K[i] - 1
		cnt := mp[S[k>>t]] + t
		k -= (k >> t) << t
		cnt += bits.OnesCount(uint(k))
		cnt %= 3
		fmt.Fprintf(_w, "%c\n", rune('A'+cnt))
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
