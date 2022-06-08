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

func main() {
	defer _w.Flush()
	_s.Split(bufio.ScanWords)
	_s.Buffer([]byte{}, math.MaxInt32)
	N, M := ScanInt(), ScanInt()
	A := make([]int, N+1)
	C := make([]int, N+M+1)
	for i := 0; i <= N; i++ {
		A[i] = ScanInt()
	}
	for i := 0; i <= N+M; i++ {
		C[i] = ScanInt()
	}
	B := Solve(N, M, A, C)
	for _, v := range B {
		fmt.Fprintf(_w, "%d ", v)
	}
	fmt.Fprintln(_w)
}

func Solve(N, M int, A, C []int) []int {
	B := make([]int, M+1)
	for i := M; i >= 0; i-- {
		B[i] = C[i+N]
		for j := 1; i+j <= M; j++ {
			if N-j<0{
				continue
			}
			B[i] -= B[i+j] * A[N-j]
		}
		B[i] /= A[N]
	}
	return B
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

var _s, _w = bufio.NewScanner(os.Stdin), bufio.NewWriter(os.Stdout)
