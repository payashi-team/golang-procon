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

func main() {
	defer _w.Flush()
	var N int
	fmt.Fscan(_r, &N)
	C := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(_r, &C[i])
	}
	ans := Solve(N, C)
	fmt.Fprintf(_w, "%d\n", ans)
}

func Solve(N int, C []int) int {
	cnts := make([]int, 0)
	cnt := 1
	cur := -1
	C = append(C, 2)
	for i := 0; i <= N; i++ {
		if C[i] == cur {
			cnt++
		} else if cur != -1 {
			cnts = append(cnts, cnt)
			cnt = 1
		}
		cur = C[i]
	}
	if len(cnts) == 1 {
		return -1
	}
	if C[0] == C[N-1] {
		cnts[0] += cnts[len(cnts)-1]
		cnts = cnts[:len(cnts)-1]
	}
	return (MaxInt(cnts...) + 1) / 2
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
