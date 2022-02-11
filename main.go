package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const (
	INF = int(1 << 60)
	// MOD = int(1e9 + 7)
	MOD = 998244353
)

func main() {
	defer _w.Flush()
	var N, M int
	fmt.Fscan(_r, &N, &M)
	H := make([]int, N)
	W := make([]int, M)
	for i := 0; i < N; i++ {
		fmt.Fscan(_r, &H[i])
	}
	for i := 0; i < M; i++ {
		fmt.Fscan(_r, &W[i])
	}
	ans := Solve(N, M, H, W)
	fmt.Fprintf(_w, "%d\n", ans)
}

func Solve(N, M int, H, W []int) int {
	sort.Ints(H)
	SumEvens := make([]int, N/2+1) // (0, 1) +... (2*k+2, 2*k+3)
	for i := 1; 2*i-1 < N; i++ {
		SumEvens[i] = SumEvens[i-1] + H[2*i-1] - H[2*i-2]
	}
	SumOdds := make([]int, N/2+1) // (1, 2) + ... (2*k-1, 2*k)
	for i := 1; 2*i < N; i++ {
		SumOdds[i] = SumOdds[i-1] + H[2*i] - H[2*i-1]
	}
	ret := INF
	for _, w := range W {
		idx := sort.SearchInts(H, w)
		diff := SumEvens[idx/2] + SumOdds[N/2] - SumOdds[idx/2]
		if idx&1 == 0 {
			diff += H[idx] - w
		} else {
			diff += w - H[idx-1]
		}
		ret = MinInt(ret, diff)
	}
	return ret
}

func MinInt(nums ...int) int {
	ret := INF
	for _, v := range nums {
		if ret > v {
			ret = v
		}
	}
	return ret
}

var _r, _w = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
