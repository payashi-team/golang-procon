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
	var N, a int
	fmt.Fscan(_r, &a, &N)
	ans := Solve(a, N)
	fmt.Fprintf(_w, "%d\n", ans)
}

type Item struct {
	cnt, num int
}

func Solve(a, N int) int {
	que := make([]Item, 0)
	memo := make(map[int]int)
	que = append(que, Item{0, 1})
	for len(que) > 0 {
		p := que[0]
		que = que[1:]
		if p.num > N*10 || (p.num != 1 && memo[p.num] != 0) {
			continue
		}
		memo[p.num] = p.cnt
		if p.num == N {
			return p.cnt
		}
		que = append(que, Item{p.cnt + 1, p.num * a})
		if p.num >= 10 && p.num%10 != 0 {
			S := strconv.Itoa(p.num)
			S = S[len(S)-1:] + S[:len(S)-1]
			num, _ := strconv.Atoi(S)
			que = append(que, Item{p.cnt + 1, num})
		}
	}
	return -1
}

var _r, _w = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
