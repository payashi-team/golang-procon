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
	st := Point{ni() - 1, ni() - 1}
	W := nl()
	field := make([][]int, 9)
	for i := 0; i < 9; i++ {
		field[i] = make([]int, 9)
		nums := nl()
		for j, c := range nums {
			field[i][j] = int(c - '0')
		}
	}
	ans := Solve(st, W, field)
	fmt.Fprintf(wr, "%s\n", ans)
}

func Solve(st Point, W string, field [][]int) string {
	xfield := make([][]int, 3*9-2)
	for i := 0; i < 3*9-2; i++ {
		xfield[i] = make([]int, 3*9-2)
		var i_ int
		if i < 8 {
			i_ = 8 - i
		} else if i < 2*9-1 {
			i_ = i - 8
		} else {
			i_ = 24 - i
		}
		for j := 0; j < 3*9-2; j++ {
			var j_ int
			if j < 8 {
				j_ = 8 - j
			} else if j < 2*9-1 {
				j_ = j - 8
			} else {
				j_ = 24 - j
			}
			xfield[i][j] = field[i_][j_]
		}
	}
	st.x += 8
	st.y += 8
	var d Point
	switch W {
	case "R":
		d = Point{0, 1}
	case "L":
		d = Point{0, -1}
	case "U":
		d = Point{-1, 0}
	case "D":
		d = Point{1, 0}
	case "RU":
		d = Point{-1, 1}
	case "RD":
		d = Point{1, 1}
	case "LU":
		d = Point{-1, -1}
	case "LD":
		d = Point{1, -1}
	}
	ret := make([]byte, 4)
	for i := 0; i < 4; i++ {
		ret[i] = byte(xfield[st.y+i*d.x][st.x+i*d.y] + '0')
	}
	return string(ret)
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
