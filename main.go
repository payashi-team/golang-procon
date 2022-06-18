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

type Range struct {
	l, r int
}

func main() {
	defer _w.Flush()
	_s.Split(bufio.ScanWords)
	N := nextInt()
	X := make([]int, N)
	C := make([]int, N)
	for i := 0; i < N; i++ {
		X[i] = nextInt() - 1
	}
	for i := 0; i < N; i++ {
		C[i] = nextInt()
	}
	Solve(N, X, C)
}

func Solve(N int, X, C []int) {
	Xinv := make([][]int, N)
	for i := 0; i < N; i++ {
		Xinv[i] = make([]int, 0)
	}
	for i := 0; i < N; i++ {
		Xinv[X[i]] = append(Xinv[X[i]], i)
	}
	used := make([]bool, N)
	Y := make([]int, 0)
	var dfs func(int)
	dfs = func(pos int) {
		used[pos] = true
		nxt := X[pos]
		if !used[nxt] {
			dfs(nxt)
		}
		Y = append(Y, pos)
	}
	for i := 0; i < N; i++ {
		if !used[i] {
			dfs(i)
		}
	}
	// fmt.Fprintf(_w, "%v\n", Y)
	used = make([]bool, N)
	tmp := make([]int, 0)
	var dfs2 func(int)
	dfs2 = func(pos int) {
		used[pos] = true
		for _, nxt := range Xinv[pos] {
			if !used[nxt] {
				dfs2(nxt)
			}
		}
		tmp = append(tmp, pos)
	}
	ret := 0
	for i := N - 1; i >= 0; i-- {
		if !used[Y[i]] {
			tmp = make([]int, 0)
			dfs2(Y[i])
			// fmt.Fprintf(_w, "%v\n", tmp)
			if len(tmp) > 1 {
				min := INF
				for _, v := range tmp {
					if min > C[v] {
						min = C[v]
					}
				}
				ret += min
			}
		}
	}
	fmt.Fprintf(_w, "%d\n", ret)
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
