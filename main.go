package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	defer _w.Flush()
	var H, W int
	fmt.Fscan(_r, &H, &W)
	A := make([][]byte, H)
	for i := 0; i < H; i++ {
		fmt.Fscan(_r, &A[i])
	}
	ans := Solve(H, W, A)
	if ans {
		fmt.Fprintf(_w, "Yes\n")
	} else {
		fmt.Fprintf(_w, "No\n")
	}
}

func Solve(H, W int, A [][]byte) bool {
	cnt := make(map[byte]int)
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			cnt[A[i][j]]++
		}
	}
	var p, q int
	for _, v := range cnt {
		p += v / 4
		if v%2 == 1 {
			q++
		}
	}
	if H%2+W%2 == 0 {
		return p == H*W/4
	} else if H*W%2 == 1 {
		return p >= (H-1)*(W-1)/4 && q == 1
	} else {
		if H%2 == 0 {
			H, W = W, H
		}
		return p >= (H-1)*W/4 && q == 0
	}
}

var _r, _w = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
