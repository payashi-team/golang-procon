package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

const (
	INF = int(1 << 61)
	MOD = int(1e9 + 7)
	// MOD = 998244353
)

func main() {
	defer _w.Flush()
	var N, M int
	fmt.Fscan(_r, &N, &M)
	Solve(N, M)
}

func Solve(N, M int) {
	if N == 1 {
		if M != 0 {
			fmt.Fprintf(_w, "-1\n")
			return
		} else {
			fmt.Fprintf(_w, "1 2\n")
			return
		}
	}
	if M > N-2 || M < 0 {
		fmt.Fprintf(_w, "-1\n")
		return
	}
	k := M + 1
	fmt.Printf("1 %d\n", 4*M+4)
	for i := 0; i <= M; i++ {
		fmt.Printf("%d %d\n", 4*i+2, 4*i+3)
	}
	for j := 0; j < N-(k+1); j++ {
		fmt.Printf("%d %d\n", 4*j+2+4*M+8, 4*j+3+4*M+8)
	}
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

var _r, _w = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
