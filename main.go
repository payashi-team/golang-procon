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

type Card struct {
	mark  rune
	num   string
	loyal bool
}

func main() {
	defer _w.Flush()
	var S string
	fmt.Fscan(_r, &S)
	cards := make([]Card, 0)
	loyals := []string{"A", "10", "J", "Q", "K"}
	for pos := 0; pos < len(S); pos++ {
		var c Card
		c.mark = rune(S[pos])
		pos++
		if S[pos] == '1' {
			c.num = "10"
			pos++
		} else {
			c.num = S[pos : pos+1]
		}
		for _, l := range loyals {
			if c.num == l {
				c.loyal = true
			}
		}
		cards = append(cards, c)
	}
	ans := Solve(cards)
	str := ""
	for _, c := range ans {
		str += string(c.mark) + c.num
	}
	if str == "" {
		str = "0"
	}
	fmt.Fprintf(_w, "%s\n", str)
}

func Solve(cards []Card) []Card {
	// for _, c := range cards {
	// 	fmt.Printf("%c@%s(%v) ", c.mark, c.num, c.loyal)
	// }
	cnt := make(map[rune]int)
	trash := make([]Card, 0)
	var mark rune
	for i, c := range cards {
		if !c.loyal {
			continue
		}
		cnt[c.mark]++
		if cnt[c.mark] >= 5 {
			// fmt.Fprintf(_w, "%c is collected!\n", c.mark)
			mark = c.mark
			cards = cards[:i]
			break
		}
	}
	for _, c := range cards {
		if !c.loyal || c.mark != mark {
			trash = append(trash, c)
		}
	}
	return trash
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
