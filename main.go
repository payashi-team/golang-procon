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
	var N int
	fmt.Fscan(_r, &N)
	ans := Solve(N)
	fmt.Fprintf(_w, "%s", ans)
}

func Solve(N int) string {
	Row := func(pos int) string {
		ret := ""
		for i := 0; i < N; i++ {
			if (i-pos+N)%N < 3 {
				ret += "#"
			} else {
				ret += "."
			}
		}
		ret += "\n"
		return ret
	}
	ret := ""
	if N%3 == 0 {
		for i := 0; i < 3; i++ {
			for j := 0; j < N/3; j++ {
				ret += Row(j * 3)
			}
		}
		return ret
	}
	arr := make([]int, N)
	pos := 0
	for i := 0; i < N; i++ {
		if pos == N-1 {
			arr[i] = N - 2
		} else if pos == N-2 {
			arr[i] = N - 1
		} else {
			arr[i] = pos
		}
		pos = (pos + 3) % N
	}
	for _, v := range arr {
		ret += Row(v)
	}
	return ret
}

func AbsInt(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

func MaxInt(nums ...int) int {
	ret := -1
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

