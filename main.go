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
	x, y int
}

func main() {
	defer _w.Flush()

	var H, W, N, M int
	fmt.Fscan(_r, &H, &W, &N, &M)
	lights := make([]Point, N)
	blocks := make([]Point, M)
	for i := 0; i < N; i++ {
		fmt.Fscan(_r, &lights[i].x, &lights[i].y)
	}
	for i := 0; i < M; i++ {
		fmt.Fscan(_r, &blocks[i].x, &blocks[i].y)
	}
	ans := Solve(H, W, N, M, lights, blocks)
	fmt.Fprintf(_w, "%d\n", ans)
}

func Solve(H, W, N, M int, lights, blocks []Point) int {
	// 0: undefined, 1: light, 2: blocks, 3: lit(horizontally), 4: lit(vertically)
	lit := make([][]byte, H)
	for i := 0; i < H; i++ {
		lit[i] = make([]byte, W)
	}
	// init lit
	for i := 0; i < N; i++ {
		p := &lights[i]
		p.x--
		p.y--
		lit[p.x][p.y] = 1
	}
	for i := 0; i < M; i++ {
		p := &blocks[i]
		p.x--
		p.y--
		lit[p.x][p.y] = 2
	}
	// simulate
	for _, p := range lights {
		for i := p.x - 1; 0 <= i && i < H; i-- {
			if lit[i][p.y] == 0 {
				lit[i][p.y] = 3
			} else {
				break
			}
		}
		for i := p.x + 1; 0 <= i && i < H; i++ {
			if lit[i][p.y] == 0 {
				lit[i][p.y] = 3
			} else {
				break
			}
		}
	}
	for _, p := range lights {
		for i := p.y - 1; 0 <= i && i < W; i-- {
			if lit[p.x][i] == 0 || lit[p.x][i] == 3 {
				lit[p.x][i] = 4
			} else {
				break
			}
		}
		for i := p.y + 1; 0 <= i && i < W; i++ {
			if lit[p.x][i] == 0 || lit[p.x][i] == 3 {
				lit[p.x][i] = 4
			} else {
				break
			}
		}
	}
	ret := 0
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			if lit[i][j]&1 == 1 || lit[i][j] == 4 {
				ret++
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

var _r, _w = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
