package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

// const (
// INF = int(1 << 62)
// MOD = int(1e9 + 7)
// MOD = 998244353
// )

type Cord struct {
	x, y float64
}

func (c Cord) Norm() float64 {
	return math.Sqrt(c.x*c.x + c.y*c.y)
}

func (c Cord) Plus(d Cord) Cord {
	return Cord{c.x + d.x, c.y + d.y}
}

func (c Cord) Scalar(alpha float64) Cord {
	return Cord{c.x * alpha, c.y * alpha}
}

func main() {
	defer _w.Flush()
	var N int
	var P0, PM Cord
	fmt.Fscan(_r, &N)
	fmt.Fscan(_r, &P0.x, &P0.y)
	fmt.Fscan(_r, &PM.x, &PM.y)
	ans := Solve(N, P0, PM)
	fmt.Fprintf(_w, "%.8f %.8f\n", ans.x, ans.y)
}

func Solve(N int, P0, PM Cord) Cord {
	C := P0.Plus(PM).Scalar(0.5)
	CP0 := P0.Plus(C.Scalar(-1))
	r := CP0.Norm()
	theta0 := math.Atan2(CP0.y, CP0.x)
	Ry, Rx := math.Sincos(theta0 + 2*math.Pi/float64(N))
	// fmt.Printf("C: %v\nCP0: %v\nr: %f\ntheta0: %f\n", C, CP0, r, theta0)
	return C.Plus(Cord{Rx, Ry}.Scalar(r))
}

var _r, _w = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
