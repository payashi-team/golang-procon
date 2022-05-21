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
	var S string
	fmt.Fscan(_r, &N, &S)
	ans := Solve(N, S)
	fmt.Fprintf(_w, "%d\n", ans)
}

func Solve(N int, S string) int {
	pairs := make([][]byte, 0)
	chars := []byte{'A', 'B', 'X', 'Y'}
	for _, c1 := range chars {
		for _, c2 := range chars {
			pairs = append(pairs, []byte{c1, c2})
		}
	}
	count := func(p1, p2 []byte) int {
		cnt := 0
		for i := 0; i+2 <= N; i++ {
			if S[i] == p1[0] && S[i+1] == p1[1] {
				cnt++
				i++
			} else if S[i] == p2[0] && S[i+1] == p2[1] {
				cnt++
				i++
			}
		}
		return N - cnt
	}
	ans := INF
	M := len(pairs)
	for i := 0; i < M; i++ {
		for j := i + 1; j < M; j++ {
			ans = MinInt(ans, count(pairs[i], pairs[j]))
		}
	}
	return ans
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
