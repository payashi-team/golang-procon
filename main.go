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

var sc = bufio.NewScanner(os.Stdin)
var wr = bufio.NewWriter(os.Stdout)

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)

	A, B, K := NextInt(), NextInt(), NextInt()

	p, q := A+B, B
	num := 1 // pCq
	for i := 0; i < q; i++ {
		num *= p - i
		num /= i + 1
	}
	ret := ""
	for i := 0; i < A+B; i++ {
		tmp := num * (p - q) / p
		// use a
		if tmp >= K {
			ret += "a"
			num = tmp
			// use b
		} else {
			ret += "b"
			q--
			num = num - tmp
			K -= tmp
		}
		p--
	}
	fmt.Fprintf(wr, "%s\n", ret)
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
