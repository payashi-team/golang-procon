package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	INF = int(1 << 62)
	// MOD = int(1e9 + 7)
	MOD = 998244353
)

func main() {
	defer _w.Flush()
	var H, W int
	fmt.Fscan(_r, &H, &W)
	S := make([]string, H)
	for i := 0; i < H; i++ {
		fmt.Fscan(_r, &S[i])
	}
	ans := Solve(H, W, S)
	fmt.Fprintf(_w, "%d\n", ans)
}

func Solve(H, W int, S []string) int {
	ret := 1
	for k := 0; k <= H+W-2; k++ {
		var r, b, t int
		for x := 0; x <= k; x++ {
			y := k - x
			if y >= H || x >= W {
				continue
			}
			if S[y][x] == 'R' {
				r++
			} else if S[y][x] == 'B' {
				b++
			} else {
				t++
			}
		}
		if r*b > 0 {
			return 0
		} else if r+b == 0 {
			ret *= 2
			ret %= MOD
		}
	}
	return ret
}

var _r, _w = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
