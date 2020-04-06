package main

import "fmt"

var coins =  map[int]int {
	1:3,
	5:2,
	10:1,
	50:3,
	100:0,
	500:2,
}

var t = []int {1,5,10,50,100,500}

var a = 620

func main() {
	ans := 0
	for i:=5; i>0; i-- {
		num := min(a / t[i], coins[t[i]])
		a -= t[i] * numN
		ans += num
	}

	fmt.Print(ans)
}

func min(a,b int) int {
	r := a
	if a > b {
		r = b
	}

	return r
}