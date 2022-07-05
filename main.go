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
	x, y int
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	N, K := ni(), ni()
	points := make([]Point, N)
	for i := 0; i < N; i++ {
		points[i] = Point{ni(), ni()}
	}
	ans := Solve(N, K, points)
	if ans < 0 {
		fmt.Fprintf(wr, "Infinity\n")
	} else {
		fmt.Fprintf(wr, "%d\n", ans)
	}
}

func Solve(N, K int, ps []Point) int {
	if K == 1 {
		return -1
	}
	used := make([][]bool, N)
	for i := 0; i < N; i++ {
		used[i] = make([]bool, N)
	}
	ret := 0
	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			if used[i][j] {
				continue
			}
			qs := []int{i, j}
			for k := j + 1; k < N; k++ {
				p1 := Subtract(ps[j], ps[i])
				p2 := Subtract(ps[k], ps[i])
				if p1.x*p2.y == p2.x*p1.y {
					qs = append(qs, k)
				}
			}
			cnt := len(qs)
			if cnt >= K {
				ret++
			}
			for ii := 0; ii < cnt; ii++ {
				for jj := ii + 1; jj < cnt; jj++ {
					qi, qj := qs[ii], qs[jj]
					if qi > qj {
						qi, qj = qj, qi
					}
					used[qi][qj] = true
				}
			}
		}
	}
	return ret
}

func Subtract(a, b Point) Point {
	return Point{a.x - b.x, a.y - b.y}

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
