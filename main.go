package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

const (
	INF = int(1 << 61)
	MOD = int(1e9 + 7)
	// MOD = 998244353
)

var sc = bufio.NewScanner(os.Stdin)
var wr = bufio.NewWriter(os.Stdout)

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	N, M, X, Y := ni(), ni(), ni(), ni()
	A := make([]int, N)
	B := make([]int, M)
	for i := 0; i < N; i++ {
		A[i] = ni()
	}
	for i := 0; i < M; i++ {
		B[i] = ni()
	}
	ans := Solve(N, M, X, Y, A, B)
	fmt.Fprintf(wr, "%d\n", ans)

}

func Solve(N, M, X, Y int, A, B []int) int {
	pos := 0
	ret := 0
	for {
		aidx := sort.Search(N, func(i int) bool { return pos <= A[i] })
		if aidx == N {
			break
		}
		pos = A[aidx] + X
		bidx := sort.Search(M, func(i int) bool { return pos <= B[i] })
		if bidx == M {
			break
		}
		pos = B[bidx] + Y
		ret++
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
