package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	defer _w.Flush()
	var S string
	fmt.Fscan(_r, &S)
	ans := Solve(S)
	fmt.Fprintf(_w, "%d\n", ans)
}

func Solve(S string) int {
	N:=len(S)
	ans:=(N-1)*N
	for i := 0; i < N; i++ {
		if S[i]=='U'{
			ans+=i
		}else{
			ans+=N-1-i
		}
	}
	return ans
}

var _r, _w = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
