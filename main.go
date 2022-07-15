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
	X := nl()
	ans := Solve(X)
	fmt.Fprintf(wr, "%s\n", ans)
}

func Solve(X string) string {
	K := len(X)

	for a0 := 1; a0 < 10; a0++ {
		for d := -9; d < 10; d++ {
			ok := true
			more := false
			for i := 0; i < K; i++ {
				t := int(X[i] - '0')
				a := a0 + i*d
				if a < 0 || 10 <= a || (!more && a < t) {
					ok = false
					break
				}
				if t < a {
					more = true
				}
			}
			if ok {
				ans := make([]byte, K)
				for i := 0; i < K; i++ {
					ans[i] = byte('0' + a0 + i*d)
				}
				return string(ans)
			}
		}
	}
	return "ng"
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
