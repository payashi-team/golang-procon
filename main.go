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

// var _r = bufio.NewReader(os.Stdin)
var _w = bufio.NewWriter(os.Stdout)
var _s = bufio.NewScanner(os.Stdin)

func main() {
	// defer _w.Flush()
	_s.Split(bufio.ScanWords)
	N := nextInt()
	from2 := make([]int, N+1)
	for i := 3; i <= N; i++ {
		fmt.Printf("? %d %d\n", 2, i)
		from2[i] = nextInt()
	}
	alphas := make([]int, 0)
	betas := make([]int, 0)
	for i := 3; i <= N; i++ {
		if from2[i] == 1 {
			alphas = append(alphas, i)
		} else if from2[i] == 2 {
			betas = append(betas, i)
		}
	}
	if len(alphas) == 0 {
		fmt.Printf("! 1\n")
		return
	}
	mindAlpha := INF // min distance from 1 to alphas
	mindBeta := INF  // min distance from 1 to betas
	alpha := -1
	beta := -1
	for _, a := range alphas {
		fmt.Printf("? %d %d\n", 1, a)
		d := nextInt()
		if mindAlpha > d {
			mindAlpha = d
			alpha = a
		}
	}
	if mindAlpha != 2 {
		fmt.Printf("! %d\n", mindAlpha+1)
		return
	}
	for _, b := range betas {
		fmt.Printf("? %d %d\n", 1, b)
		d := nextInt()
		if mindBeta > d {
			mindBeta = d
			beta = b
		}
	}
	if mindBeta != 1 {
		fmt.Printf("! 1\n")
		return
	}
	fmt.Printf("? %d %d\n", alpha, beta)
	d := nextInt()
	if d == 1 {
		fmt.Printf("! 3\n")
	} else {
		fmt.Printf("! 1\n")
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

func nextInt() int {
	_s.Scan()
	i, e := strconv.Atoi(_s.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func nextLine() string {
	_s.Scan()
	return _s.Text()
}
