package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

const (
	INF = int(1 << 61)
	// MOD = int(1e9 + 7)
	MOD = 998244353
)

func main() {
	defer _w.Flush()
	var N, Q int
	fmt.Fscan(_r, &N, &Q)
	nums := make([]int, N+1)
	for i := 0; i < Q; i++ {
		var L, R int
		fmt.Fscan(_r, &L, &R)
		nums[L-1]++
		nums[R]--
	}
	for i := 0; i < N; i++ {
		nums[i+1] += nums[i]
	}
	nums = nums[:N]
	for _, v := range nums {
		fmt.Printf("%d", v&1)
	}
	fmt.Println()
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

var _r, _w = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
