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
	X := make([]int, N)
	Y := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(_r, &X[i], &Y[i])
	}
	ans := Solve(N, X, Y)
	fmt.Fprintf(_w, "%d\n", ans)
}

func Solve(N int, X, Y []int) int {
	ranges := make([][]int, 2)
	for i := 0; i < 2; i++ {
		ranges[i] = make([]int, 2)
		ranges[i][0] = INF
		ranges[i][1] = -INF
	}
	for i := 0; i < N; i++ {
		x := X[i]
		y := Y[i]
		for j := 0; j < 2; j++ {
			val := x
			if j&1 == 0 {
				val += y
				// fmt.Printf("%d + %d = %d\n", x, y, val)
			} else {
				val -= y
				// fmt.Printf("%d - %d = %d\n", x, y, val)
			}
			ranges[j][0] = MinInt(ranges[j][0], val)
			ranges[j][1] = MaxInt(ranges[j][1], val)
		}
	}
	ans := -1
	for i := 0; i < 2; i++ {
		ans = MaxInt(ans, ranges[i][1]-ranges[i][0])
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

var _r, _w = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
