package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	INF = int(1 << 60)
	// MOD = int(1e9 + 7)
	MOD = 998244353
)

func main() {
	defer _w.Flush()
	var N int
	fmt.Fscan(_r, &N)
	ans := Solve(N)
	if ans {
		fmt.Fprintf(_w, "Yes\n")
	} else {
		fmt.Fprintf(_w, "No\n")
	}
}

func Solve(N int) bool {
	const M = int(1 << 31)
	return -M <= N && N < M
}

var _r, _w = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
