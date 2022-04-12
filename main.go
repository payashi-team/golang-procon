package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

const (
	INF = int(1 << 61)
	MOD = int(1e9 + 7)
	// MOD = 998244353
)

func main() {
	defer _w.Flush()
	S := make([]string, 3)
	for i := 0; i < 3; i++ {
		fmt.Fscan(_r, &S[i])
	}
	Solve(S)
}

func Solve(S []string) {
	set := make(map[rune]struct{})
	for i := 0; i < 3; i++ {
		for _, c := range S[i] {
			set[c] = struct{}{}
		}
	}
	if len(set) > 10 {
		fmt.Fprintln(_w, "UNSOLVABLE")
		return
	}
	arr := make([]rune, 0)
	for k, _ := range set {
		arr = append(arr, k)
	}
	mp := make(map[rune]int)
	used := make(map[int]bool)
	for _, c := range arr {
		mp[c] = -1
	}
	convert := func(T string) int {
		ret := 0
		for i := 0; i < len(T); i++ {
			ret *= 10
			ret += mp[rune(T[i])]
		}
		return ret
	}
	check := func() bool {
		for i := 0; i < 3; i++ {
			if mp[rune(S[i][0])] == 0 {
				return false
			}
		}
		return convert(S[0])+convert(S[1]) == convert(S[2])
	}
	var dfs func(int) bool
	dfs = func(num int) bool {
		if num == len(arr) {
			if check() {
				for i := 0; i < 3; i++ {
					fmt.Printf("%d\n", convert(S[i]))
				}
				return true
			} else {
				return false
			}
		}
		c := arr[num]
		for i := 0; i < 10; i++ {
			if used[i] {
				continue
			}
			mp[c] = i
			used[i] = true
			if dfs(num + 1) {
				return true
			}
			mp[c] = -1
			used[i] = false
		}
		return false
	}
	if !dfs(0) {
		fmt.Fprintln(_w, "UNSOLVABLE")
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

var _r, _w = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
