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
	MOD = int(1e9 + 7)
	// MOD = 998244353
)

var sc = bufio.NewScanner(os.Stdin)
var wr = bufio.NewWriter(os.Stdout)

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	N := NextInt()
	A := make([]int, N)
	for i := 0; i < N; i++ {
		A[i] = NextInt()
	}
	ans := Solve(N, A)
	fmt.Fprintf(wr, "%d\n", ans)
}

func Solve(N int, A []int) int {
	mp := make(map[int]int)
	for _, v := range A {
		mp[v]++
	}
	ret := 0
	M := MaxInt(A...)
	// i*j = k
	for i := 1; i <= M; i++ {
		for j := 1; j <= M/i; j++ {
			ret += mp[i] * mp[j] * mp[i*j]
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

func NextInt() int {
	sc.Scan()
	x, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return x
}

func NextLine() string {
	sc.Scan()
	return sc.Text()
}
