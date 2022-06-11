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
	// MOD = 998244353
	MOD = 10007
)

func main() {
	defer _w.Flush()
	_s.Split(bufio.ScanWords)
	N := nextInt()
	P := make([]int, N)
	I := make([]int, N)
	for i := 0; i < N; i++ {
		P[i] = nextInt() - 1
	}
	for i := 0; i < N; i++ {
		I[i] = nextInt() - 1
	}
	Solve(N, P, I)
}

func Solve(N int, P, I []int) {
	mp := make(map[int]int)
	for i := 0; i < N; i++ {
		mp[I[i]] = i
	}
	for i := 0; i < N; i++ {
		P[i] = mp[P[i]]
	}
	fmt.Printf("%v\n", P)
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

func nextInt() int {
	_s.Scan()
	i, e := strconv.Atoi(_s.Text())
	if e != nil {
		panic(e)
	}
	return i
}

// func nextLine() string {
// 	_s.Scan()
// 	return _s.Text()
// }

var _s, _w, _r = bufio.NewScanner(os.Stdin), bufio.NewWriter(os.Stdout), bufio.NewReader(os.Stdout)
