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
	ans := Solve(N)
	for _, v := range ans {
		fmt.Fprintf(_w, "%d ", v)
	}
	fmt.Fprintln(_w)
}

func Solve(N int) []int {
	nums := make([]bool, 10001)
	init := []int{6, 10, 15}
	if N==3{
		return init
	}
	for _, v := range init {
		for i := v; i <= 10000; i += v {
			nums[i] = true
		}
	}
	ret := make([]int, 0)
	for i := 0; i <= 10000 && N > 0; i++ {
		if nums[i] {
			ret = append(ret, i)
			N--
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
