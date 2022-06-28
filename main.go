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

type Edge struct {
	from, to, cost int
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	N := ni()
	edges := make([][]int, N)
	for i := 0; i < N; i++ {
		edges[i] = make([]int, 0)
	}
	for i := 0; i < N-1; i++ {
		x, y := ni()-1, ni()-1
		edges[x] = append(edges[x], y)
		edges[y] = append(edges[y], x)
	}
	Q := ni()
	questions := make([]Query, Q)
	for i := 0; i < Q; i++ {
		questions[i] = Query{ni() - 1, ni() - 1}
	}
	Solve(N, edges, Q, questions)
}

type Query struct {
	a, b int
}

func Solve(N int, edges [][]int, Q int, questions []Query) {
	K := 1
	for 1<<K < N {
		K++
	}
	dist := make([]int, N)
	for i := 0; i < N; i++ {
		dist[i] = -1
	}
	parent := make([][]int, K)
	for i := 0; i < K; i++ {
		parent[i] = make([]int, N)
	}
	var dfs func(int, int, int)
	dfs = func(u, p, d int) {
		dist[u] = d
		parent[0][u] = p
		for _, v := range edges[u] {
			if v == p {
				continue
			}
			dfs(v, u, d+1)
		}
	}
	dfs(0, -1, 0)
	for k := 0; k < K-1; k++ {
		for u := 0; u < N; u++ {
			if parent[k][u] == -1 {
				parent[k+1][u] = -1
			} else {
				parent[k+1][u] = parent[k][parent[k][u]]
			}
		}
	}
	for _, q := range questions {
		u, v := q.a, q.b
		// v is deeper
		if dist[u] > dist[v] {
			u, v = v, u
		}
		for k := 0; k < K; k++ {
			if (dist[v]-dist[u])>>k&1 == 1 {
				v = parent[k][v]
			}
		}
		if u == v {
			fmt.Fprintf(wr, "%d\n", dist[q.a]+dist[q.b]-2*dist[u]+1)
			continue
		}
		for k := K - 1; k >= 0; k-- {
			if parent[k][u] != parent[k][v] {
				u = parent[k][u]
				v = parent[k][v]
			}
		}
		w := parent[0][u]
		fmt.Fprintf(wr, "%d\n", dist[q.a]+dist[q.b]-2*dist[w]+1)
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
