package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	defer _w.Flush()
	var A, B, C int64
	fmt.Fscan(_r, &A, &B, &C)
	ans := Solve(A, B, C)
	fmt.Fprintf(_w, "%s\n", ans)
}

func Solve(a, b, c int64) string {
	if c < a+b {
		return "No"
	}
	ans := 4*a*b < (c-(a+b))*(c-(a+b))
	if ans {
		return "Yes"
	} else {
		return "No"
	}
}

var _r, _w = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
