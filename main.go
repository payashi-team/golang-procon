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
	fmt.Printf("%d\n", Solve(N))
}

func Solve(N int) int {
	prime := make([]bool, int(1e6))
	for i := 2; i < int(1e6); i++ {
		prime[i] = true
	}
	for i := 2; i < int(1e6); i++ {
		if !prime[i] {
			continue
		}
		for j := i * 2; j < int(1e6); j += i {
			prime[j] = false
		}
	}
	sum := make([]int, int(1e6))
	for i := 1; i < int(1e6); i++ {
		if prime[i] {
			sum[i] = sum[i-1] + 1
		} else {
			sum[i] = sum[i-1]
		}
	}
	ret := 0
	for q := 2; q < int(1e6); q++ {
		if !prime[q] {
			continue
		}
		p := N / (q * q * q)
		p = MinInt(p, q-1)
		if p == 0 {
			break
		}
		// fmt.Printf("p: %d, q: %d\n", p, q)
		ret += sum[p]
	}
	return ret
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

