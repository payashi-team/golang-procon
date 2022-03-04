package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

const (
	INF = int(1 << 61)
	MOD = int(1e9 + 7)
	// MOD = 998244353
)

type Query struct {
	command string
	n, m    int
}

func main() {
	defer _w.Flush()
	var S string
	fmt.Fscan(_r, &S)
	ans:=Solve(S)
	fmt.Printf("%d\n", ans)
}

func Solve(S string) int {
	cnt := make(map[rune]int)
	for _, c := range S {
		cnt[c]++
	}
	var odd, even int
	for _, v := range cnt {
		if v%2 == 1 {
			odd++
			even += v - 1
		} else {
			even += v
		}
	}
	if odd == 0 {
		return len(S)
	} else {
		return (even/2)/odd*2 + 1
	}
}

func AbsInt(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

func MaxInt(nums ...int) int {
	ret := -1
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

var _r, _w = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
