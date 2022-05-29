package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

const (
	INF = int(1 << 61)
	// MOD = int(1e9 + 7)
	MOD = 998244353
)

func main() {
	defer _w.Flush()
	_s.Split(bufio.ScanWords)
	_s.Buffer([]byte{}, math.MaxInt32)
	N, K := ScanInt(), ScanInt()
	R := make([]int, N)
	for i := 0; i < N; i++ {
		R[i] = ScanInt()
	}
	ans := Solve(N, K, R)
	fmt.Fprintf(_w, "%.8f\n", ans)
}

func Solve(N, K int, R []int) float64 {
	sort.Slice(R, func(i, j int) bool { return R[i] > R[j] })
	ans := .0
	for k := 1; k <= K; k++ {
		ans += float64(R[k-1]) * math.Pow(0.5, float64(k))
	}
	return ans
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

func ScanInt() int {
	_s.Scan()
	num, err := strconv.Atoi(_s.Text())
	if err != nil {
		panic(err)
	}
	return num
}

var _s, _w = bufio.NewScanner(os.Stdin), bufio.NewWriter(os.Stdout)
