package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

const (
	INF = int(1 << 61)
	// MOD = int(1e9 + 7)
	MOD = 998244353
)

func main() {
	defer _w.Flush()
	var N int
	fmt.Fscanf(_r, "%d\n", &N)
	C := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(_r, &C[i])
	}
	edges := make([][]int, N)
	for i := 0; i < N; i++ {
		edges[i] = make([]int, 0)
	}
	for i := 0; i < N-1; i++ {
		var a, b int
		fmt.Fscan(_r, &a, &b)
		a--
		b--
		edges[a] = append(edges[a], b)
		edges[b] = append(edges[b], a)
	}
	ans := Solve(N, C, edges)
	for _, v := range ans {
		fmt.Fprintf(_w, "%d\n", v)
	}
}

func Solve(N int, C []int, edges [][]int) []int {
	ret := make([]int, 0)
	used := make([]bool, N)
	var dfs func(int, map[int]int)
	dfs = func(u int, mp map[int]int) {
		if mp[C[u]] == 0 {
			ret = append(ret, u+1)
		}
		mp[C[u]]++
		used[u] = true
		for _, v := range edges[u] {
			if used[v] {
				continue
			}
			dfs(v, mp)
		}
		mp[C[u]]--
	}
	dfs(0, make(map[int]int))
	sort.Ints(ret)
	return ret
}

// Both a and b are sorted
func MinDist(a, b []int) int {
	ret := INF
	for _, v := range a {
		idx := sort.Search(len(b), func(i int) bool { return v <= b[i] })
		if idx == 0 {
			ret = MinInt(ret, b[idx]-v)
		} else if idx == len(b) {
			ret = MinInt(ret, v-b[idx-1])
		} else {
			ret = MinInt(ret, b[idx]-v, v-b[idx-1])
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

var _r, _w = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
