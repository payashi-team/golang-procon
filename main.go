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

type Edge struct {
	to, color, cost int
}

type Query struct {
	color, cost, u, v int
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	N := ni()
	Q := ni()
	edges := make([][]Edge, N)
	for i := 0; i < N; i++ {
		edges[i] = make([]Edge, 0)
	}
	for i := 0; i < N-1; i++ {
		a, b := ni()-1, ni()-1
		c, d := ni()-1, ni()
		edges[a] = append(edges[a], Edge{b, c, d})
		edges[b] = append(edges[b], Edge{a, c, d})
	}
	qs := make([]Query, Q)
	for i := 0; i < Q; i++ {
		qs[i] = Query{ni() - 1, ni(), ni() - 1, ni() - 1}
	}
	ans := Solve(N, Q, edges, qs)
	for _, v := range ans {
		fmt.Fprintf(wr, "%d\n", v)
	}
}

type Memo struct {
	qi, color, cost, delta int
}

func Solve(N, Q int, edges [][]Edge, qs []Query) []int {
	// lowest common ancestor
	t := 1
	K := 0
	for t < N {
		t *= 2
		K++
	}
	depth := make([]int, N)
	parent := make([][]int, N)
	for i := 0; i < N; i++ {
		parent[i] = make([]int, K)
		parent[i][0] = -1
	}
	var lca func(int, int, int)
	lca = func(cur, par, dep int) {
		parent[cur][0] = par
		depth[cur] = dep
		for _, e := range edges[cur] {
			if e.to == par {
				continue
			}
			lca(e.to, cur, dep+1)
		}
	}
	lca(0, -1, 0)
	for i := 0; i < K-1; i++ {
		for j := 0; j < N; j++ {
			if parent[j][i] == -1 {
				parent[j][i+1] = -1
			} else {
				parent[j][i+1] = parent[parent[j][i]][i]
			}
		}
	}
	getlca := func(a, b int) int {
		if depth[a] > depth[b] {
			a, b = b, a
		}
		diff := depth[b] - depth[a]
		// b is deeper
		for i := 0; i < K; i++ {
			if (diff>>i)&1 == 1 {
				b = parent[b][i]
			}
		}
		if a == b {
			return a
		}
		for i := K - 1; i >= 0; i-- {
			if parent[a][i] != parent[b][i] {
				a = parent[a][i]
				b = parent[b][i]
			}
		}
		return parent[a][0]
	}

	memo := make([][]Memo, N)
	for i := 0; i < N; i++ {
		memo[i] = make([]Memo, 0)
	}
	for i, q := range qs {
		memo[q.u] = append(memo[q.u], Memo{i, q.color, q.cost, 1})
		memo[q.v] = append(memo[q.v], Memo{i, q.color, q.cost, 1})
		w := getlca(q.u, q.v)
		if w != -1 {
			memo[w] = append(memo[w], Memo{i, q.color, q.cost, -2})
		}
	}
	ans := make([]int, Q)
	cnt := make([]int, N-1)  // for each color
	dist := make([]int, N-1) // for each color
	var dfs func(int, int, int)
	dfs = func(cur, par, sum int) {
		for _, m := range memo[cur] {
			ans[m.qi] += (sum - dist[m.color] + cnt[m.color]*m.cost) * m.delta
		}
		for _, e := range edges[cur] {
			if e.to == par {
				continue
			}
			cnt[e.color]++
			dist[e.color] += e.cost
			dfs(e.to, cur, sum+e.cost)
			cnt[e.color]--
			dist[e.color] -= e.cost
		}
	}
	dfs(0, -1, 0)
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
