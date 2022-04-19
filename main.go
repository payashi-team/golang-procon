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

type Point struct {
	x, y float64
}

func main() {
	defer _w.Flush()
	var p0 Point
	fmt.Fscan(_r, &p0.x, &p0.y)
	var N int
	fmt.Fscan(_r, &N)
	points := make([]Point, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(_r, &points[i].x, &points[i].y)
	}
	ans := Solve(p0, N, points)
	fmt.Fprintf(_w, "%.10f\n", ans)
}

func Solve(p0 Point, N int, points []Point) float64 {
	ans := 1000.0
	for i := 0; i < N; i++ {
		p1 := points[i]
		p2 := points[(i+1)%N]
		s := -Dot(Minus(p1, p0), Minus(p2, p1)) / Dot(Minus(p2, p1), Minus(p2, p1))
		if s < 0 || 1 < s {
			continue
		}
		p3 := Minus(Plus(Scale(p1, 1-s), Scale(p2, s)), p0)
		dist := math.Sqrt(Dot(p3, p3))
		if ans > dist {
			ans = dist
		}
	}
	return ans
}

func Dot(p1, p2 Point) float64 {
	return p1.x*p2.x + p1.y*p2.y
}

func Scale(p Point, alpha float64) Point {
	return Point{p.x * alpha, p.y * alpha}
}

func Plus(p1, p2 Point) Point {
	return Point{p1.x + p2.x, p1.y + p2.y}
}

func Minus(p1, p2 Point) Point {
	return Point{p1.x - p2.x, p1.y - p2.y}
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
