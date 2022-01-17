package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)
	seats := make([]string, N)
	fmt.Printf("0\n")
	fmt.Scan(&seats[0])
	if seats[0] == "Vacant" {
		return
	}
	l := 0
	r := N
	for {
		mid := (l + r) / 2
		fmt.Printf("%d\n", mid)
		fmt.Scan(&seats[mid])
		if seats[mid] == "Vacant" {
			break
		}
		if (seats[l] == seats[mid]) == ((mid-l)%2 == 1) {
			r = mid
		} else {
			l = mid
		}
	}
}

// var _r, _w = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
