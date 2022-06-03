package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

const (
	INF = int(1 << 61)
	// MOD = int(1e9 + 7)
	MOD = 998244353
)

func main() {
	defer _w.Flush()
	_s.Split(bufio.ScanWords)
	_s.Buffer([]byte{}, math.MaxInt32)
	N, K := ScanInt(), ScanInt()
	A := make([]int, N)
	for i := 0; i < N; i++ {
		A[i] = ScanInt()
	}
	ans := Solve(N, K, A)
	fmt.Fprintf(_w, "%d\n", ans)
}

type Item struct {
	val, idx int
}

func Solve(N, K int, A []int) int {
	items := make([]Item, N)
	for i := 0; i < N; i++ {
		items[i] = Item{A[i], i}
	}
	sort.Slice(items, func(i, j int) bool {
		if items[i].val != items[j].val {
			return items[i].val < items[j].val
		} else {
			return items[i].idx > items[j].idx
		}
	})
	pos := -1
	ret := INF
	for i := 0; i < N; i++ {
		if items[i].idx >= K {
			if pos == -1 {
				continue
			}
			ret = MinInt(ret, items[i].idx-pos)
		} else {
			pos = MaxInt(pos, items[i].idx)
		}
	}
	if ret == INF {
		return -1
	} else {
		return ret
	}

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

func ScanInt() int {
	_s.Scan()
	num, err := strconv.Atoi(_s.Text())
	if err != nil {
		panic(err)
	}
	return num
}

var _s, _w = bufio.NewScanner(os.Stdin), bufio.NewWriter(os.Stdout)
