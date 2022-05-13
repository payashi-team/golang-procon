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

func main() {
	defer _w.Flush()
	var H, W int
	fmt.Fscan(_r, &H, &W)
	field := make([][]byte, H)
	for i := 0; i < H; i++ {
		fmt.Fscan(_r, &field[i])
	}
	Solve(H, W, field)
}

func Solve(H, W int, field [][]byte) {
	X := make([][]bool, H)
	Y := make([][]bool, H)
	for i := 0; i < H; i++ {
		X[i] = make([]bool, W)
		Y[i] = make([]bool, W)
	}
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			black := true
			for dx := -1; dx <= 1; dx++ {
				for dy := -1; dy <= 1; dy++ {
					nx := i + dx
					ny := j + dy
					if nx < 0 || H <= nx || ny < 0 || W <= ny {
						continue
					}
					if field[nx][ny] != '#' {
						black = false
					}
				}
			}
			if black {
				X[i][j] = true
			}
		}
	}
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			if !X[i][j] {
				continue
			}
			Y[i][j] = true
			for dx := -1; dx <= 1; dx++ {
				for dy := -1; dy <= 1; dy++ {
					nx := i + dx
					ny := j + dy
					if nx < 0 || H <= nx || ny < 0 || W <= ny {
						continue
					}
					Y[nx][ny] = true
				}
			}
		}
	}
	ok:=true
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			if (field[i][j]=='#')!=Y[i][j]{
				ok = false
			}
		}
	}
	if !ok{
		fmt.Fprintf(_w, "impossible\n")
	}else{
		fmt.Fprintf(_w, "possible\n")
		for i := 0; i < H; i++ {
			for j := 0; j < W; j++ {
				if X[i][j]{
					fmt.Fprint(_w, "#")
				}else{
					fmt.Fprint(_w, ".")
				}
			}
			fmt.Fprintln(_w)
			
		}
	}

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
