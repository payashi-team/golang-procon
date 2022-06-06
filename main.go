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

func main() {
	defer _w.Flush()
	_s.Split(bufio.ScanWords)
	_s.Buffer([]byte{}, math.MaxInt32)
	N := ScanInt()
	bits := make([]bool, N)
	for i := 0; i < N; i++ {
		if ScanInt() == 1 {
			bits[i] = true
		}
	}
	ans := Solve(N, bits)
	fmt.Fprintf(_w, "%d\n", ans)
}

func Solve(N int, bits []bool) int {
	nums := make([]int, N+1)
	for i := 0; i < N; i++ {
		if bits[i] {
			nums[i+1] = nums[i] - 1
		} else {
			nums[i+1] = nums[i] + 1
		}
	}
	maxs := make([]int, N+1)
	mins := make([]int, N+1)
	for i := 0; i < N; i++ {
		maxs[i+1] = MaxInt(maxs[i], nums[i])
		mins[i+1] = MinInt(mins[i], nums[i])
	}
	var max, min int
	for i := 0; i < N; i++ {
		max = MaxInt(max, nums[i+1]-mins[i])
		min = MinInt(min, nums[i+1]-maxs[i])
	}
	return max - min + 1
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
