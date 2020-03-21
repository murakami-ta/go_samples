package main

import (
	"fmt"
)

var n = 4
var a = []int{1,2,4,7}
var k = 15

func main() {
	if dps(0, 0) {
		fmt.Print("Yes")
	} else {
		fmt.Print("No")
	}
}

func dps(i int, sum int) bool {
	if sum == k {
		return true
	}

	if i >= n {
		return false
	}

	if dps(i+1, sum + a[i]) {
		return true
	}

	if dps(i+1, sum) {
		return true
	}

	return false
}
