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
	A := make([][]int, 2*N)
	for i := 0; i < 2*N; i++ {
		A[i] = make([]int, 2*N)
	}
	for i := 0; i < 2*N-1; i++ {
		for j := i + 1; j < 2*N; j++ {
			fmt.Fscan(_r, &A[i][j])
			A[j][i] = A[i][j]
		}
	}
	ans := Solve(N, A)
	fmt.Fprintf(_w, "%d\n", ans)
}

func Solve(N int, A [][]int) int {
	var dfs func(int, []bool) int
	dfs = func(score int, used []bool) int {
		pos := -1
		for i, v := range used {
			if !v {
				pos = i
				break
			}
		}
		if pos < 0 {
			// fmt.Printf("used: %v, score: %d\n", used, score)
			return score
		}
		used[pos] = true
		ret := 0
		for i := 0; i < 2*N; i++ {
			if used[i] {
				continue
			}
			used[i] = true
			ret = MaxInt(ret, dfs(score^A[pos][i], used))
			used[i] = false
		}
		used[pos] = false
		return ret
	}
	used := make([]bool, 2*N)
	return dfs(0, used)
}

func MaxInt(nums ...int) int {
	ret := 0
	for _, v := range nums {
		if ret < v {
			ret = v
		}
	}
	return ret
}

var _r, _w = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
