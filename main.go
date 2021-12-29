package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	defer _w.Flush()
	var S string
	fmt.Fscan(_r, &S)
	ans := Solve(S)
	fmt.Fprintf(_w, "%s\n", ans)
}

func Solve(S string) string {
	var a, b, c int
	for i := 0; i < len(S); i++ {
		if S[i]=='a'{
			a++;
		}else if S[i]=='b'{
			b++;
		}else{
			c++;
		}
	}
	max:=MaxInt([]int{a, b, c})
	min:=-MaxInt([]int{-a, -b, -c})
	if max-min>1{
		return "NO"
	}else{
		return "YES"
	}
}

func MaxInt(nums []int)int{
	ret:=-int(1e5)
	for _, v := range nums {
		if ret<v{
			ret = v
		}
	}
	return ret
}

var _r, _w = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
