package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	defer _w.Flush()
	var K, A, B int
	fmt.Fscan(_r, &K, &A, &B)
	ans := Solve(K, A, B)
	fmt.Fprintf(_w, "%d\n", ans)
}

func Solve(K, A, B int) int {
	M := K - A + 1
	if B-A <= 2 || M < 0 {
		return K + 1
	} else {
		return A + M/2*(B-A) + M%2
	}
}

var _r, _w = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
