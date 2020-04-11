package main

import (
	"fmt"
	"sort"
)

var n = 5
var s = []int{1,2,4,6,8}
var t = []int{3,5,7,9,10}

func main() {
	var pairs []pair
	for i:=0; i<n; i++ {
		pairs = append(pairs, pair{s[i], t[i]})
	}

	sort.Slice(pairs, func(i, j int) bool { return pairs[i].t < pairs[j].t })

	c := 0
	l := 0
	for i:=0; i<n; i++ {
		if l <= pairs[i].s {
			c++
			l = pairs[i].t
		}
	}

	fmt.Print(c)
}


type pair struct {
	s int
	t int
}