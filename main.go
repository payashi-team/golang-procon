package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
)

const (
	INF = int(1 << 61)
	// MOD = int(1e9 + 7)
	MOD = 998244353
)

func main() {
	defer _w.Flush()
	var N int
	fmt.Fscan(_r, &N)
	P := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(_r, &P[i])
		P[i]--
	}
	ans, err := Solve(N, P)
	if err != nil {
		fmt.Fprintf(_w, "-1\n")
	} else {
		for i := 0; i < N-1; i++ {
			fmt.Fprintf(_w, "%d\n", ans[i])
		}
	}
}

func Solve(N int, P []int) ([]int, error) {
	Q := make(map[int]int, N)
	used := make([]bool, N-1)
	ret := make([]int, 0)
	for i, v := range P {
		Q[v] = i
	}
	for i := 0; i < N; i++ {
		for Q[i] > i {
			pos := Q[i]
			if used[pos-1] {
				return ret, errors.New("wtf")
			} else {
				used[pos-1] = true
				ret = append(ret, pos)
			}
			// swap
			j := P[pos-1]
			P[pos], P[pos-1] = P[pos-1], P[pos]
			Q[i], Q[j] = Q[j], Q[i]
		}
	}
	if len(ret) != N-1 {
		return ret, errors.New("no bueno")
	}
	return ret, nil
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

var _r, _w = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
