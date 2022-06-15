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
	MOD = int(1e9 + 7)
	// MOD = 998244353
)

// var _r = bufio.NewReader(os.Stdin)
var _w = bufio.NewWriter(os.Stdout)
var _s = bufio.NewScanner(os.Stdin)

func main() {
	defer _w.Flush()
	_s.Split(bufio.ScanWords)
	N := nextInt()
	T := make([]int, N)
	for i := 0; i < N; i++ {
		T[i] = nextInt()
	}
	sort.Ints(T)
	cnt := make(map[int]int)
	ret := 0
	for i := 0; i < N; i++ {
		ret += T[i] * (N - i)
		cnt[T[i]]++
	}
	fact := make([]int, 10001)
	fact[0] = 1
	for i := 1; i <= 10000; i++ {
		fact[i] = fact[i-1] * i
		fact[i] %= MOD
	}
	ret2 := 1
	for _, v := range cnt {
		ret2 *= fact[v]
		ret2 %= MOD
	}
	fmt.Fprintf(_w, "%d\n%d\n", ret, ret2)

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

func nextLine() string {
	_s.Scan()
	return _s.Text()
}
