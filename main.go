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
	var N, Q int
	fmt.Fscan(_r, &N, &Q)
	X := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(_r, &X[i])
	}
	A := make([]int, N-1)
	B := make([]int, N-1)
	for i := 0; i < N-1; i++ {
		fmt.Fscan(_r, &A[i], &B[i])
	}
	V := make([]int, Q)
	K := make([]int, Q)
	for i := 0; i < Q; i++ {
		fmt.Fscan(_r, &V[i], &K[i])
	}
	ans := Solve(N, X, A, B, Q, V, K)
	for _, v := range ans {
		fmt.Fprintf(_w, "%d\n", v)
	}
}

func Solve(N int, X, A, B []int, Q int, V, K []int) []int {
	MAX_K := 20
	edges := make([][]int, N)
	for i := 0; i < N-1; i++ {
		a := A[i] - 1
		b := B[i] - 1
		edges[a] = append(edges[a], b)
		edges[b] = append(edges[b], a)
	}
	called := make([]bool, N)
	biggest := make([][]int, N) // 各頂点について、大きな順に20個いれる
	var dfs func(int) []int
	dfs = func(v int) []int {
		called[v] = true
		biggest[v] = []int{X[v]}
		for _, u := range edges[v] {
			if called[u] {
				continue
			}
			biggest[v] = append(biggest[v], dfs(u)...)
		}
		sort.Slice(biggest[v], func(i, j int) bool { return biggest[v][i] > biggest[v][j] })
		if len(biggest[v]) > MAX_K {
			biggest[v] = biggest[v][:MAX_K]
		}
		return biggest[v]
	}
	dfs(0)
	// fmt.Printf("%v\n", biggest)
	ans := make([]int, Q)
	for i := 0; i < Q; i++ {
		v := V[i] - 1
		k := K[i] - 1
		ans[i] = biggest[v][k]
	}
	return ans
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
