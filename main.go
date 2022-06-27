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

type Road struct {
	x, y, z int
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	N, M := ni(), ni()
	dist := make([][]int, N)
	for i := 0; i < N; i++ {
		dist[i] = make([]int, N)
		for j := 0; j < N; j++ {
			dist[i][j] = INF
		}
		dist[i][i] = 0
	}
	for i := 0; i < M; i++ {
		A, B, C := ni(), ni(), ni()
		A--
		B--
		dist[A][B] = C
		dist[B][A] = C
	}
	K := ni()
	roads := make([]Road, K)
	for i := 0; i < K; i++ {
		x, y, z := ni(), ni(), ni()
		x--
		y--
		roads[i] = Road{x, y, z}
	}
	ans := Solve(N, M, dist, K, roads)
	for i := 0; i < K; i++ {
		fmt.Fprintf(wr, "%d\n", ans[i])
	}
}

func Solve(N, M int, dist [][]int, K int, roads []Road) []int {
	for k := 0; k < N; k++ {
		for i := 0; i < N; i++ {
			for j := 0; j < N; j++ {
				dist[i][j] = MinInt(dist[i][j], dist[i][k]+dist[k][j])
			}
		}
	}
	sum := 0
	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			sum += dist[i][j]
		}
	}
	ret := make([]int, 0)
	for _, r := range roads {
		diff := dist[r.x][r.y] - r.z
		if diff <= 0 {
			ret = append(ret, sum)
			continue
		}
		memo := make([]Pair, 0)
		for i := 0; i < N; i++ {
			for j := i + 1; j < N; j++ {
				ndist := MinInt(dist[i][r.x]+dist[r.y][j], dist[i][r.y]+dist[r.x][j]) + r.z
				if ndist < dist[i][j] {
					memo = append(memo, Pair{i, j, ndist})
					sum -= dist[i][j] - ndist
				}
			}
		}
		for _, m := range memo {
			dist[m.x][m.y] = m.dist
			dist[m.y][m.x] = m.dist
		}
		ret = append(ret, sum)
	}
	return ret
}

type Pair struct {
	x, y, dist int
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
