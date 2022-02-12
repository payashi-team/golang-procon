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
	var S, T []byte
	fmt.Fscan(_r, &S, &T)
	ans := Solve(N, S, T)
	fmt.Fprintf(_w, "%d\n", ans)
}

func Solve(N int, S, T []byte) int {
	for i := 1; i < N; i++ {
		if (S[i-1] == '0') == (S[i] == '1') {
			S[i] = '1'
		} else {
			S[i] = '0'
		}
		if (T[i-1] == '0') == (T[i] == '1') {
			T[i] = '1'
		} else {
			T[i] = '0'
		}
	}
	ret := 0
	for i, j := 0, 0; i < N; i++ {
		if i > j {
			j = i
		}
		if S[j] == T[i] {
			continue
		}
		for j+1 < N && S[j] == S[j+1] {
			j++
		}
		if j+1 == N {
			return -1
		}
		j++
		ret += j - i
	}
	return ret
}

var _r, _w = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
