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
	N, D, K := ni(), ni(), ni()
	L := make([]int, D)
	R := make([]int, D)
	S := make([]int, K)
	T := make([]int, K)
	for i := 0; i < D; i++ {
		L[i], R[i] = ni()-1, ni()-1
	}
	for i := 0; i < K; i++ {
		S[i], T[i] = ni()-1, ni()-1
	}
	ans := Solve(N, D, K, L, R, S, T)
	for _, v := range ans {
		fmt.Fprintf(wr, "%d\n", v)
	}
}

func Solve(N, D, K int, L, R, S, T []int) []int {
	ret := make([]int, K)
	for i := 0; i < D; i++ {
		for j := 0; j < K; j++ {
			if ret[j] > 0 {
				continue
			}
			if L[i] <= S[j] && S[j] <= R[i] {
				if L[i] <= T[j] && T[j] <= R[i] {
					ret[j] = i + 1
					S[j] = T[j]
				} else if S[j] < T[j] {
					S[j] = R[i]
				} else {
					S[j] = L[i]
				}
			}
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
