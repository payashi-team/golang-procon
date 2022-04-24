package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

const (
	INF = int(1 << 61)
	// MOD = int(1e9 + 7)
	MOD = 998244353
)

func main() {
	// fmt.Printf("%d\n", Gcd(24, 36))
	// var p, q int
	// ExtGcd(1, 1, &p, &q)
	// fmt.Printf("\n%d*%d+%d*%d=%d\n", 1, p, 1, q, 1)
	defer _w.Flush()
	var T int
	fmt.Fscan(_r, &T)
	for i := 0; i < T; i++ {
		var N, A, B, X, Y, Z int
		fmt.Fscan(_r, &N, &A, &B, &X, &Y, &Z)
		ans := Solve(N, A, B, X, Y, Z)
		fmt.Fprintf(_w, "%d\n", ans)
	}
}

func Solve(N, A, B, X, Y, Z int) int {
	if A > B {
		A, B = B, A
		Y, Z = Z, Y
	}
	usea := true
	useb := true
	if A*X <= Y {
		usea = false
	}
	if B*X <= Z {
		useb = false
	}
	mosta := (N/A)*Y + (N%A)*X
	mostb := (N/B)*Z + (N%B)*X
	if !usea && !useb {
		return N * X
	} else if usea && !useb {
		return mosta
	} else if !usea && useb {
		return mostb
	} else {
		ans := MinInt(mosta, mostb)
		for i := 0; i < 100000; i++ {
			numB := N/B - i
			if numB < 0 {
				break
			}
			numA := (N - numB) % A
			numOne := N - numA - numB
			cost := numB*Z + numA*Y + numOne*X
			fmt.Printf("cost: %d\n", cost)
			ans = MinInt(ans, cost)
		}
		return ans
	}
}

// a*x+b*y=1
func ExtGcd(a, b int, x, y *int) int {
	d := a
	if b != 0 {
		d = ExtGcd(b, a%b, y, x)
		*y -= (a / b) * (*x)
	} else {
		*x = 1
		*y = 0
	}
	return d
}

func Gcd(a, b int) int {
	if a > b {
		a, b = b, a
	}
	if a == 0 {
		return b
	}
	return Gcd(b%a, a)
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
