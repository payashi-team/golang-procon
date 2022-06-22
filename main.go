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

var sc = bufio.NewScanner(os.Stdin)
var wr = bufio.NewWriter(os.Stdout)

type Range struct {
	l, r int
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)

	N, D := NextInt(), NextInt()
	X := make([]Range, N)
	for i := 0; i < N; i++ {
		X[i].l, X[i].r = NextInt(), NextInt()
	}
	sort.Slice(X, func(i, j int) bool {
		if X[i].r != X[j].r {
			return X[i].r < X[j].r
		} else {
			return X[i].l < X[j].l
		}
	})
	pos := -1
	ret := 0
	for i := 0; i < N; i++ {
		if X[i].l <= pos {
			continue
		} else {
			ret++
			pos = X[i].r + D - 1
		}
	}
	fmt.Fprintf(wr, "%d\n", ret)
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

func NextInt() int {
	sc.Scan()
	x, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return x
}

func NextLine() string {
	sc.Scan()
	return sc.Text()
}
