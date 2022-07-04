package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

const (
	INF = int(1 << 61)
	// MOD = int(1e9 + 7)
	MOD = 998244353
)

var sc = bufio.NewScanner(os.Stdin)
var wr = bufio.NewWriter(os.Stdout)

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	N, K := ni(), ni()
	A := make([]int, N)
	for i := 0; i < N; i++ {
		A[i] = ni()
	}
	ans := Solve(N, K, A)
	for _, v := range ans {
		fmt.Fprintf(wr, "%d ", v)
	}
	fmt.Fprintln(wr)
}

func Solve(N, K int, A []int) []int {
	for k := 0; k < K; k++ {
		B := make([]int, N)
		for i := 0; i < N; i++ {
			l := i - A[i]
			r := i + A[i] + 1
			B[MaxInt(0, l)]++
			if r < N {
				B[r]--
			}
		}
		for i := 0; i < N-1; i++ {
			B[i+1] += B[i]
		}
		update := false
		for i := 0; i < N; i++ {
			A[i] = B[i]
			if A[i] < N {
				update = true
			}
		}
		if !update {
			break
		}
	}
	return A
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

func ni() int {
	sc.Scan()
	x, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return x
}

func nl() string {
	sc.Scan()
	return sc.Text()
}
