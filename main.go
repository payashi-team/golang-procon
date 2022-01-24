package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// const (
// MOD = int(1e9 + 7)
// MOD = 998244353
// )

func main() {
	defer _w.Flush()
	var K int
	fmt.Fscan(_r, &K)
	var S, T string
	fmt.Fscan(_r, &S, &T)
	ans := Solve(K, S, T)
	fmt.Fprintf(_w, "%.8f\n", ans)
}

func Solve(K int, S, T string) float64 {
	ret := 0.
	remain := make(map[int]int)
	for i := 1; i < 10; i++ {
		remain[i] = K
	}
	for i := 0; i < 4; i++ {
		num, _ := strconv.Atoi(S[i : i+1])
		remain[num]--
		num, _ = strconv.Atoi(T[i : i+1])
		remain[num]--
	}
	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			var prob int
			if i == j {
				prob = remain[i] * (remain[i] - 1)
			} else {
				prob = remain[i] * remain[j]
			}
			if Score(S, i) > Score(T, j) {
				ret += float64(prob)
			}
		}
	}
	ret /= float64(K*9-8) * float64(K*9-9)
	return ret
}

func Score(S string, p int) int {
	S = S[:4] + strconv.Itoa(p)
	cnt := make(map[int]int)
	for i := 0; i < 5; i++ {
		num, _ := strconv.Atoi(S[i : i+1])
		cnt[num]++
	}
	ret := 0
	for i := 1; i < 10; i++ {
		t := i
		for j := 0; j < cnt[i]; j++ {
			t *= 10
		}
		ret += t
	}
	return ret
}

var _r, _w = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
