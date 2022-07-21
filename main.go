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
	R := make([]int, N)
	for i := 0; i < N; i++ {
		R[i] = ni()
	}
	ans := Solve(N, R)
	fmt.Fprintf(wr, "%d\n", ans)
}

func Solve(N int, R []int) int {
	dp := make([][]int, 2)
	for i := 0; i < 2; i++ {
		dp[i] = make([]int, N)
	}
	// dp0 increasing
	// dp1 decreasing
	for i := 0; i < N; i++ {
		dp[0][i] = 1
		dp[1][i] = 1
		for j := 0; j < i; j++ {
			if R[i] > R[j] {
				dp[0][i] = MaxInt(dp[0][i], dp[1][j]+1)
			} else if R[i] < R[j] {
				dp[1][i] = MaxInt(dp[1][i], dp[0][j]+1)
			}
		}
	}
	ret := MaxInt(dp[0][N-1], dp[1][N-1])
	if ret < 3 {
		return 0
	} else {
		return ret
	}
}

type UFind struct {
	dep, par []int
}

func NewUFind(N int) *UFind {
	uf := new(UFind)
	uf.par = make([]int, N)
	uf.dep = make([]int, N)
	for i := 0; i < N; i++ {
		uf.par[i] = i
	}
	return uf
}

func (uf *UFind) Root(x int) int {
	if uf.par[x] == x {
		return x
	}
	uf.par[x] = uf.Root(uf.par[x])
	return uf.par[x]
}

func (uf *UFind) Same(x, y int) bool {
	return uf.Root(x) == uf.Root(y)
}

func (uf *UFind) Unite(x, y int) {
	x = uf.Root(x)
	y = uf.Root(y)
	if x == y {
		return
	}
	if uf.dep[x] == uf.dep[y] {
		uf.dep[x]++
	} else if uf.dep[x] < uf.dep[y] {
		x, y = y, x
	}
	uf.par[y] = x
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
