package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
  defer _w.Flush()
  var N int
  fmt.Fscan(_r, &N)
  s:=make([]bool, N)
  for i := 0; i < N; i++ {
    var op string
    fmt.Fscan(_r, &op)
    if(op=="AND"){
      s[i]=true
    }
  }
  ans:=Solve(N, s)
  fmt.Fprintf(_w, "%d\n", ans)
}

func Solve(N int, s []bool)int{
  t:=make([]int, N+1)
  f:=make([]int, N+1)
  t[0] = 1
  f[0] = 1
  for i := 0; i < N; i++ {
    if(s[i]){
      t[i+1] = t[i]
      f[i+1] = t[i]+2*f[i]
    }else{
      t[i+1] = 2*t[i] + f[i]
      f[i+1] = f[i]
    }
  }
  return t[N]
}

var _r, _w = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)