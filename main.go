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
	B := make([]int, N)
	for i := 1; i < N; i++ {
		B[i] = ScanInt() - 1
	}
	ans := Solve(N, B)
	fmt.Fprintf(_w, "%d\n", ans)
}

func Solve(N int, B []int) int {
	edges := make([][]int, N)
	for i := 0; i < N; i++ {
		edges[i] = make([]int, 0)
	}
	for i := 1; i < N; i++ {
		edges[B[i]] = append(edges[B[i]], i)
	}
	var dfs func(int) int
	dfs = func(pos int) int {
		if len(edges[pos]) == 0 {
			return 1
		}
		min := INF
		max := 0
		for _, v := range edges[pos] {
			pay := dfs(v)
			min = MinInt(min, pay)
			max = MaxInt(max, pay)
		}
		return min + max + 1
	}
	return dfs(0)
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
