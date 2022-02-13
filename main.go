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
	A := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(_r, &A[i])
	}
	ans := Solve(N, A)
	fmt.Fprintf(_w, "%d\n", ans)
}

func Solve(N int, A []int) int {
	cost := make([]int, 32)
	bitcnts := make([]int, 32)
	for _, num := range A {
		digit := 0
		for num > 0 {
			if num&1 == 1 {
				bitcnts[digit]++
			}
			digit++
			num /= 2
		}
	}
	bigness := 1
	for i := 0; i < 32; i++ {
		cost[i] = (N - 2*bitcnts[i]) * bigness
		bigness *= 2
	}
	ret := 0
	for _, num := range A {
		score := 0
		digit := 0
		for num > 0 {
			if num&1 == 1 {
				score += cost[digit]
			}
			digit++
			num /= 2
		}
		ret = MaxInt(ret, score)
	}
	for _, num := range A {
		ret += num
	}
	return ret
}

func MaxInt(nums ...int) int {
	ret := 0
	for _, v := range nums {
		if ret < v {
			ret = v
		}
	}
	return ret
}

var _r, _w = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
