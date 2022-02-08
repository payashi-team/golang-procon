package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

const (
	// INF = int(1 << 60)
	// MOD = int(1e9 + 7)
	MOD = 998244353
)

var Dirs = []string{
	"N", "NNE", "NE", "ENE", "E", "ESE", "SE", "SSE", "S", "SSW", "SW", "WSW", "W", "WNW", "NW", "NNW",
}
var Ws = []float64{
	0.2, 1.5, 3.3, 5.4, 7.9, 10.7, 13.8, 17.1, 20.7, 24.4, 28.4, 32.6,
}

func main() {
	defer _w.Flush()
	var Deg, Dis int
	fmt.Fscan(_r, &Deg, &Dis)
	Dir, W := Solve(Deg, Dis)
	fmt.Fprintf(_w, "%s %d\n", Dir, W)
}

func Solve(Deg, Dis int) (string, int) {
	Dir := Dirs[int(math.Floor((float64(Deg)*0.1+11.25)/22.5))%16]
	wind := math.Round(float64(Dis)/60.*10) / 10.
	W := 12
	for i, w := range Ws {
		if wind <= w {
			W = i
			break
		}
	}
	if W == 0 {
		Dir = "C"
	}
	return Dir, W
}

var _r, _w = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
