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
	N, C := ni(), ni()
	A := make([]int, N)
	for i := 0; i < N; i++ {
		A[i] = ni() - 1
	}
	ans := Solve(N, C, A)
	fmt.Fprintf(wr, "%d\n", ans)
}

func Solve(N, C int, A []int) int {
	ret := INF
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if i == j {
				continue
			}
			cnt := 0
			for k := 0; k < N; k++ {
				if k&1 == 0 && A[k] != i {
					cnt++
				}
				if k&1 == 1 && A[k] != j {
					cnt++
				}
			}
			ret = MinInt(ret, cnt)
		}
	}
	return ret * C
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
