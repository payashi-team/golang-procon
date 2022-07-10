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

type Point struct {
	x, y float64
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	N := ni()
	P := make([]Point, N)
	for i := 0; i < N; i++ {
		P[i].x, P[i].y = float64(ni()), float64(ni())
	}
	ans := Solve(N, P)
	fmt.Fprintf(wr, "%.9f\n", ans)
}

func Solve(N int, P []Point) float64 {
	ret := -1.0
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			dx := P[i].x - P[j].x
			dy := P[i].y - P[j].y
			dist := math.Sqrt(dx*dx + dy*dy)
			if dist > ret {
				ret = dist
			}
		}
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
