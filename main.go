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

var sc = bufio.NewScanner(os.Stdin)
var wr = bufio.NewWriter(os.Stdout)

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	sc.Scan()
	X := sc.Text()
	N := len(X)
	A := make([]int, N)
	for i := 0; i < N; i++ {
		A[i] = int(X[i] - '0')
	}
	S := make([]int, N+1) // S[i+1] = A[0] + ... + A[i]
	for i := 0; i < N; i++ {
		S[i+1] = S[i] + A[i]
	}
	T := make([]rune, 0)
	cur := 0
	for i := N; i >= 1; i-- {
		num := S[i] + cur
		cur = num / 10
		T = append(T, rune(num%10+'0'))
	}
	for cur > 0 {
		T = append(T, rune(cur%10+'0'))
		cur /= 10
	}
	M := len(T)
	for i := 0; i < M; i++ {
		fmt.Printf("%c", T[M-1-i])
	}
	fmt.Println()
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
