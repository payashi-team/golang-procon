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
	// MOD = 998244353
	MOD = 10007
)

func main() {
	defer _w.Flush()
	N := ScanInt()
	ans := Solve(N)
	fmt.Fprintf(_w, "%d\n", ans)
}

type Matrix [][]int

func NewMatrix(n, m int) Matrix {
	mtx := make(Matrix, n)
	for i := 0; i < n; i++ {
		mtx[i] = make([]int, m)
	}
	return mtx
}

// (n, m) * (m, l)
func Multiple(a, b Matrix) Matrix {
	n, m, l := len(a), len(a[0]), len(b[0])
	c := NewMatrix(n, l)
	for i := 0; i < n; i++ {
		for j := 0; j < l; j++ {
			for k := 0; k < m; k++ {
				c[i][j] += a[i][k] * b[k][j]
				c[i][j] %= MOD
			}
		}
	}
	return c
}

func Power(a Matrix, p int) Matrix {
	n := len(a)
	ret := NewMatrix(n, n)
	for i := 0; i < n; i++ {
		ret[i][i] = 1
	}
	for p > 0 {
		if p&1 == 1 {
			ret = Multiple(ret, a)
		}
		a = Multiple(a, a)
		p >>= 1
	}
	return ret
}

func Solve(N int) int {
	A := [][]int{
		{0, 1, 0},
		{0, 0, 1},
		{1, 1, 1},
	}
	A = Power(A, N-1)
	return A[0][2]

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

func ScanInt() int {
	_s.Scan()
	num, err := strconv.Atoi(_s.Text())
	if err != nil {
		panic(err)
	}
	return num
}

var _s, _w, _r = bufio.NewScanner(os.Stdin), bufio.NewWriter(os.Stdout), bufio.NewReader(os.Stdout)
