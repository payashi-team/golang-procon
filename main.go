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
	MOD = int(1e9 + 7)
	// MOD = 998244353
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
	R, C := NextInt(), NextInt()
	var s, g Point
	s.y, s.x, g.y, g.x = NextInt(), NextInt(), NextInt(), NextInt()
	s.y--
	s.x--
	g.y--
	g.x--
	field := make([][]byte, R)
	for i := 0; i < R; i++ {
		field[i] = []byte(NextLine())
	}
	ans := Solve(R, C, s, g, field)
	fmt.Fprintf(wr, "%d\n", ans)
}

func Solve(R, C int, s, g Point, field [][]byte) int {
	que := make([]Item, 0)
	que = append(que, Item{s, 0})
	used := make(map[Point]bool)
	used[s] = true
	ds := []Point{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	for len(que) > 0 {
		item := que[0]
		que = que[1:]
		if item.p == g {
			return item.dist
		}
		for _, d := range ds {
			nxt := Point{item.p.x + d.x, item.p.y + d.y}
			if !used[nxt] && 0 <= nxt.y && nxt.y < R && 0 <= nxt.x && nxt.x < C && field[nxt.y][nxt.x] == '.' {
				used[nxt] = true
				que = append(que, Item{nxt, item.dist + 1})
			}
		}
	}
	return -1

}

type Item struct {
	p    Point
	dist int
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

func NextInt() int {
	sc.Scan()
	x, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return x
}

func NextLine() string {
	sc.Scan()
	return sc.Text()
}
