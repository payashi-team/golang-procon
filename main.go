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

// var _r = bufio.NewReader(os.Stdin)
var _w = bufio.NewWriter(os.Stdout)
var _s = bufio.NewScanner(os.Stdin)

func main() {
	defer _w.Flush()
	_s.Split(bufio.ScanWords)
	A, B, C := nextInt(), nextInt(), nextInt()
	if A < B {
		A, B = B, A
	}
	if A < C {
		A, C = C, A
	}
	R := A + B + C
	r := MaxInt(A-(B+C), 0)
	ret := float64(R*R-r*r) * math.Pi
	fmt.Fprintf(_w, "%.10f\n", ret)
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
