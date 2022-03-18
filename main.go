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
	var N, M int
	fmt.Fscan(_r, &N, &M)
	A := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(_r, &A[i])
	}
	B := make([]int, M)
	C := make([]int, M)
	for i := 0; i < M; i++ {
		var c int
		fmt.Fscan(_r, &B[i], &c)
		for j := 0; j < c; j++ {
			var I int
			fmt.Fscan(_r, &I)
			I--
			C[i] |= 1 << I
		}
	}
	ans := Solve(N, M, A, B, C)
	fmt.Fprintf(_w, "%d\n", ans)
}

func Solve(N, M int, A, B, C []int) int {
	bit := 1<<9 - 1
	max := -1
	for bit < 1<<N {
		// fmt.Printf("%b\n", bit)
		score := 0
		for i := 0; i < N; i++ {
			if (bit>>i)&1 == 1 {
				score += A[i]
			}
		}
		for i := 0; i < M; i++ {
			if bits.OnesCount(uint(bit&C[i])) >= 3 {
				score += B[i]
			}
		}
		max = MaxInt(max, score)
		x := bit & -bit
		y := bit + x
		bit = ((bit &^ y) / x >> 1) | y
	}
	return max
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
