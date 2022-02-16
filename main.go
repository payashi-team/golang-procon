package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	INF = int(1 << 60)
	// MOD = int(1e9 + 7)
	MOD = 998244353
)

func main() {
	defer _w.Flush()
	var N int
	fmt.Fscan(_r, &N)
	A := make([]int, N+1)
	for i := 0; i <= N; i++ {
		fmt.Fscan(_r, &A[i])
	}
	ans := Solve(N, A)
	fmt.Fprintf(_w, "%d\n", ans)
}

func Solve(N int, A []int) int {
	curDep := 0
	maxSlot := 2
	curSlot := maxSlot
	ret := 1
	for i := 0; i <= N; i++ {
		a := A[i]
		for a > 0 {
			cnt := MinInt(curSlot, a)
			fmt.Printf("(%d - %d) x%d\n", i, curDep, cnt)
			ret += cnt * (i - curDep)
			a -= cnt
			curSlot -= cnt
			if curSlot == 0 {
				curDep++
				maxSlot *= 2
				curSlot = maxSlot
			}
		}
	}
	return ret
}

func MinInt(nums ...int) int {
	ret := INF
	for _, v := range nums {
		if ret > v {
			ret = v
		}
	}
	return ret
}

var _r, _w = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
