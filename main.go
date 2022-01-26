package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// const (
// MOD = int(1e9 + 7)
// MOD = 998244353
// )

func main() {
	defer _w.Flush()
	var N, K, M, R int
	fmt.Fscan(_r, &N, &K, &M, &R)
	S := make([]int, N)
	for i := 0; i < N-1; i++ {
		fmt.Fscan(_r, &S[i])
	}
	ans := Solve(N, K, M, R, S)
	fmt.Fprintf(_w, "%d\n", ans)
}

func Solve(N, K, M, R int, S []int) int {
	sort.Slice(S, func(i, j int) bool { return S[i] > S[j] })
	if N == K {
		sum := 0
		for _, v := range S {
			sum += v
		}
		if sum >= K*R {
			return 0
		} else if sum+M < K*R {
			return -1
		} else {
			return K*R - sum
		}
	}
	sum:=0
	for i := 0; i < K; i++ {
		sum+=S[i]
	}
	if sum>=K*R{
		return 0
	}else{
		X:=K*R-sum+S[K-1]
		if X>M{
			return -1
		}else{
			return X
		}
	}
}

var _r, _w = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
