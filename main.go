package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	// INF = int(1 << 60)
	// MOD = int(1e9 + 7)
	MOD = 998244353
)

func main() {
	defer _w.Flush()
	var N, X, M int
	fmt.Fscan(_r, &N, &X, &M)
	ans := Solve(N, X, M)
	fmt.Fprintf(_w, "%d\n", ans)
}

func Solve(N, X, M int) int {
	a := make([]int, M+1)
	cnt := make(map[int]int)
	a[1] = X
	sum := make([]int, M+1)
	for i := 1; i < M; i++ {
		sum[i] = sum[i-1] + a[i]
		cnt[a[i]] = i
		a[i+1] = a[i] * a[i] % M
		if cnt[a[i+1]] > 0 {
			s := i + 1
			t := cnt[a[i+1]]
			sum[s] = sum[s-1] + a[s]
			T := s - t
			return (sum[s]-sum[t])*((N-t+1)/T) + sum[(N-t+1)%T+(t-1)]
		}
	}
	return 0
}

var _r, _w = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
