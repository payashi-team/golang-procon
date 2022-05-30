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
	B, C := ScanInt(), ScanInt()
	ans := Solve(B, C)
	fmt.Fprintf(_w, "%d\n", ans)
}

func Solve(B, C int) int {
	var Mp, Mm int
	var mp, mm int
	if B > 0 {
		Mp = B + (C-2)/2
		Mm = -B - (C-1)/2
		mp = MaxInt(0, B-C/2)
		mm = MinInt(0, -B+(C-1)/2)
	} else {
		Mp = -B + (C-1)/2
		Mm = B - C/2
		mp = MaxInt(0, -B-(C-1)/2)
		mm = MinInt(0, B+(C-2)/2)
	}
	if mp == mm {
		return Mp - Mm + 1
	} else {
		return (Mp - mp + 1) + (mm - Mm + 1)
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

func ScanInt() int {
	_s.Scan()
	num, err := strconv.Atoi(_s.Text())
	if err != nil {
		panic(err)
	}
	return num
}

var _s, _w = bufio.NewScanner(os.Stdin), bufio.NewWriter(os.Stdout)
