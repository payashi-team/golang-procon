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

type Query struct {
	c, d int
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	N := ni()
	Q := ni()
	edges := make([][]int, N)
	for i := 0; i < N; i++ {
		edges[i] = make([]int, 0)
	}
	for i := 0; i < N-1; i++ {
		a, b := ni()-1, ni()-1
		edges[a] = append(edges[a], b)
		edges[b] = append(edges[b], a)
	}
	qs := make([]Query, Q)
	for i := 0; i < Q; i++ {
		qs[i] = Query{ni() - 1, ni() - 1}
	}
	Solve(N, Q, edges, qs)
}

func Solve(N, Q int, edges [][]int, qs []Query) {
	parent := make([][]int, N) // parent[i][j+1] := parent[parent[i][j]][j]
	K := 0
	t := 1
	for t < N {
		t *= 2
		K++
	}
	depth := make([]int, N)
	for i := 0; i < N; i++ {
		parent[i] = make([]int, K)
		parent[i][0] = -1
	}
	used := make([]bool, N)
	var dfs func(int, int)
	dfs = func(u, d int) {
		used[u] = true
		for _, v := range edges[u] {
			if used[v] {
				continue
			}
			parent[v][0] = u
			depth[v] = d + 1
			dfs(v, d+1)
		}
	}
	dfs(0, 0)
	for j := 0; j < K-1; j++ {
		for i := 0; i < N; i++ {
			if parent[i][j] == -1 {
				parent[i][j+1] = -1
			} else {
				parent[i][j+1] = parent[parent[i][j]][j]
			}
		}
	}
	for _, q := range qs {
		a, b := q.c, q.d
		if depth[a] > depth[b] {
			a, b = b, a
		}
		// b is deeper
		diff := depth[b] - depth[a] // >=0
		for i := 0; i < K; i++ {
			if (diff>>i)&1 == 1 {
				b = parent[b][i]
			}
		}
		for i := K - 1; i >= 0; i-- {
			if parent[a][i] != parent[b][i] {
				a = parent[a][i]
				b = parent[b][i]
			}
		}
		c := parent[a][0]
		dist := depth[q.c] + depth[q.d]
		if c != -1 {
			dist -= depth[c] * 2
		}
		if dist&1 == 1 {
			fmt.Println("Road")
		} else {
			fmt.Println("Town")
		}
	}
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
