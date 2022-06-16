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
	A := make([]int, N)
	mp := make(map[int]int)
	for i := 0; i < N; i++ {
		A[i] = nextInt()
		mp[A[i]]++
	}
	B := make([]int, 0)
	for k := range mp {
		B = append(B, k)
	}
	sort.Ints(B)
	C := make(map[int]int)
	for i, v := range B {
		C[v] = i
	}
	for _, v := range A {
		fmt.Fprintf(_w, "%d\n", C[v])
	}
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
