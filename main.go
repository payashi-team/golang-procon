package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

const (
	INF = int(1 << 61)
	MOD = int(1e9 + 7)
	// MOD = 998244353
)

var sc = bufio.NewScanner(os.Stdin)
var wr = bufio.NewWriter(os.Stdout)

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	H, W, A, B := ni(), ni(), ni(), ni()
	ans := Solve(H, W, A, B)
	fmt.Fprintf(wr, "%d\n", ans)

}

func Solve(H, W, A, B int) int {
	ret := 0
	var dfs func(int, int, int, int)
	dfs = func(bit, pos, a, b int) {
		if pos == H*W {
			ret++
			return
		}
		if (bit>>pos)&1 == 1 {
			dfs(bit, pos+1, a, b)
			return
		}
		bit |= 1 << pos
		// put a square tatami
		if b-1 >= 0 {
			dfs(bit, pos+1, a, b-1)
		}
		// put a long tatami horizontally
		if a-1 >= 0 && pos%W != W-1 {
			dfs(bit|1<<(pos+1), pos+2, a-1, b)
		}
		// put a long tatami vertically
		if a-1 >= 0 && pos+W < H*W {
			dfs(bit|1<<(pos+W), pos+1, a-1, b)
		}
	}
	dfs(0, 0, A, B)
	return ret
}

func Contains(x int, nums ...int) bool {
	for _, v := range nums {
		if v == x {
			return true
		}
	}
	return false
}

func AbsInt(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

func MaxInt(nums ...int) int {
	ret := -INF
	for _, v := range nums {
		if ret < v {
			ret = v
		}
	}
	return ret
}

func MinInt(nums ...int) int {
	ret := math.MaxInt64
	for _, v := range nums {
		if ret > v {
			ret = v
		}
	}
	return ret
}

func ni() int {
	sc.Scan()
	x, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return x
}

func nl() string {
	sc.Scan()
	return sc.Text()
}
