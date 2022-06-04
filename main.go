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
	// MOD = int(1e9 + 7)
	MOD = 998244353
)

func main() {
	defer _w.Flush()
	_s.Split(bufio.ScanWords)
	_s.Buffer([]byte{}, math.MaxInt32)
	N, M := ScanInt(), ScanInt()
	edges := make([][]int, N)
	for i := 0; i < N; i++ {
		edges[i] = make([]int, 0)
	}
	for i := 0; i < M; i++ {
		a, b := ScanInt()-1, ScanInt()-1
		edges[a] = append(edges[a], b)
		edges[b] = append(edges[b], a)
	}
	Q := ScanInt()
	for i := 0; i < Q; i++ {
		x, k := ScanInt()-1, ScanInt()
		que := make([]int, 0)
		que = append(que, x)
		used := []int{x}
		for t := 0; t < k; t++ {
			nextQue := make([]int, 0)
			for _, u := range que {
				for _, v := range edges[u] {
					if Contains(v, used...) {
						continue
					}
					nextQue = append(nextQue, v)
					used = append(used, v)
				}
			}
			que = nextQue
		}
		ret := 0
		for _, v := range used {
			ret += v + 1
		}
		fmt.Fprintf(_w, "%d\n", ret)
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
