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

var sc = bufio.NewScanner(os.Stdin)
var wr = bufio.NewWriter(os.Stdout)

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	N := ni()
	F := make([]int, N)
	for i := 0; i < N; i++ {
		F[i] = ni() - 1
	}
	ans := Solve(N, F)
	fmt.Fprintf(wr, "%d\n", ans)
}

func Solve(N int, F []int) int {
	E := make([]int, 2*N)
	Einv := make([][]int, 2*N)
	for i := 0; i < 2*N; i++ {
		E[i] = -1
		Einv[i] = make([]int, 0)
	}
	for i := 0; i < N; i++ {
		E[i] = N + F[i]
		E[i+N] = i
	}
	for i := 0; i < 2*N; i++ {
		v := E[i]
		if v < 0 {
			continue
		}
		Einv[v] = append(Einv[v], i)
	}
	// SCC
	cnt := 0
	used := make([]bool, 2*N)
	nums := make([]int, 0)
	var dfs func(int)
	dfs = func(u int) {
		used[u] = true
		v := E[u]
		if v >= 0 && !used[v] {
			dfs(v)
		}
		nums = append(nums, u)
	}
	for i := 0; i < 2*N; i++ {
		if !used[i] {
			dfs(i)
		}
	}
	used = make([]bool, 2*N)
	var dfs2 func(u int)
	dfs2 = func(u int) {
		used[u] = true
		for _, v := range Einv[u] {
			if used[v] {
				continue
			}
			dfs2(v)
		}
	}
	for i := 0; i < 2*N; i++ {
		u := nums[i]
		if !used[u] {
			dfs2(u)
			cnt++
		}
	}
	ret := 1
	for i := 0; i < cnt; i++ {
		ret *= 2
		ret %= MOD
	}
	ret--
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

func ni() int {
	sc.Scan()
	x, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return x
}

func nl() string {
	sc.Scan()
	return sc.Text()
}
