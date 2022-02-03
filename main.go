package main

import (
	"bufio"
	"fmt"
	"os"
)

// const (
// INF = int(1 << 62)
// MOD = int(1e9 + 7)
// MOD = 998244353
// )

func main() {
	defer _w.Flush()
	var S string
	fmt.Fscan(_r, &S)
	ans := Solve(S)
	fmt.Fprintf(_w, "%d\n", ans)
}

func Solve(S string) int {
	ret := 0
	cur := 0
	N := len(S)
	for i := 0; i < N; i++ {
		if i+2 > N {
			break
		}
		if S[i:i+2] == "25" {
			cur++
			i++
		} else {
			ret += cur * (cur + 1) / 2
			cur = 0
		}
	}
	ret += cur * (cur + 1) / 2
	return ret
}

var _r, _w = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
