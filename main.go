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
	_s.Scan()
	S := _s.Text()
	ans := Solve(S)
	fmt.Fprintf(_w, "%d\n", ans)
}

func Solve(S string) int {
	N := len(S)
	ret := 0
	cnt := make(map[byte]int)
	for i := N - 1; i >= 1; i-- {
		if S[i] == S[i-1] {
			ret += (N - 1 - i) - cnt[S[i]]
			cnt = make(map[byte]int)
			cnt[S[i]] = N - 1 - i + 2
			i--
		} else {
			cnt[S[i]]++
		}
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
