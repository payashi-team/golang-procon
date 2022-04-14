package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

const (
	INF = int(1 << 61)
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
	Solve(N, A)
}

// log2(201)<8
func Solve(N int, A []int) {
	PrintBC := func(b, c int) {
		fmt.Fprintf(_w, "Yes\n")
		B := make([]int, 0)
		bcnt := 0
		for i := 0; i < N; i++ {
			if b>>i&1 == 1 {
				B = append(B, i)
				bcnt++
			}
		}
		C := make([]int, 0)
		ccnt := 0
		for i := 0; i < N; i++ {
			if c>>i&1 == 1 {
				C = append(C, i)
				ccnt++
			}
		}
		fmt.Fprintf(_w, "%d ", bcnt)
		for _, v := range B {
			fmt.Fprintf(_w, "%d ", v+1)
		}
		fmt.Fprintln(_w)
		fmt.Fprintf(_w, "%d ", ccnt)
		for _, v := range C {
			fmt.Fprintf(_w, "%d ", v+1)
		}
		fmt.Fprintln(_w)
	}
	M := MinInt(N, 8)
	cnt := make(map[int]int)
	for bit := 1; bit < 1<<M; bit++ {
		sum := 0
		for i := 0; i < N; i++ {
			if bit>>i&1 == 1 {
				sum += A[i]
				sum %= 200
			}
		}
		if v, ok := cnt[sum]; ok {
			PrintBC(bit, v)
			return
		}
		cnt[sum] = bit
	}
	fmt.Fprintf(_w, "No\n")
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
