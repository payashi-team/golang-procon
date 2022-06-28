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
	N, M := ni(), ni()
	edges := make([]Edge, M)
	for i := 0; i < M; i++ {
		a, b, c := ni(), ni(), ni()
		a--
		b--
		edges[i] = Edge{a, b, -c}
	}
	Solve(N, M, edges)
}

func Solve(N, M int, edges []Edge) {
	// bellman-ford
	dist := make([]int, N)
	for i := 0; i < N; i++ {
		dist[i] = INF
	}
	dist[0] = 0
	for i := 0; i < N-1; i++ {
		for _, e := range edges {
			if dist[e.from] != INF && dist[e.to] > dist[e.from]+e.cost {
				dist[e.to] = dist[e.from] + e.cost
			}
		}
	}
	negative := make([]bool, N)
	for i := 0; i < N; i++ {
		for _, e := range edges {
			if dist[e.from] != INF && dist[e.to] > dist[e.from]+e.cost {
				negative[e.to] = true
			}
			if negative[e.from] {
				negative[e.to] = true
			}
		}
	}
	if negative[N-1] {
		fmt.Fprintf(wr, "inf\n")
	} else {
		fmt.Fprintf(wr, "%d\n", -dist[N-1])
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
