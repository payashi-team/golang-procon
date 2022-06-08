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

func main() {
	defer _w.Flush()
	var N, L int
	fmt.Scan(&N, &L)
	field := make([][]bool, L)
	for i := 0; i < L; i++ {
		field[i] = make([]bool, N-1)
		_s.Scan()
		S := _s.Text()
		for j := 0; j < N-1; j++ {
			field[i][j] = S[2*j+1] == '-'
		}
	}
	_s.Scan()
	T := _s.Text()
	pos := -1
	for j := 0; j < N; j++ {
		if T[2*j] == 'o' {
			pos = j
			break
		}
	}
	ans := Solve(N, L, field, pos)
	fmt.Fprintf(_w, "%d\n", ans)
}

func Solve(N, L int, field [][]bool, pos int) int {
	for i := L - 1; i >= 0; i-- {
		if pos-1 >= 0 && field[i][pos-1] {
			pos--
		} else if pos < N-1 && field[i][pos] {
			pos++
		}
	}
	return pos + 1
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

func ScanInt() int {
	_s.Scan()
	num, err := strconv.Atoi(_s.Text())
	if err != nil {
		panic(err)
	}
	return num
}

var _s, _w, _r = bufio.NewScanner(os.Stdin), bufio.NewWriter(os.Stdout), bufio.NewReader(os.Stdout)
