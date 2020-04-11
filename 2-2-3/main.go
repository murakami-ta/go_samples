package main

import "fmt"

var n = 6
var s = "ACDBCB"

func main() {
	s1 := []rune(s)
	s2 := []rune(reverse(s))
	ans := []rune(s)
	for i, i1, i2 :=0, 0, 0; i<n; i++ {
		if s1[i1] == s2[i2] {
			ans[i] = s1[i1]
			i1++
			i2++
		} else if s1[i1] > s2[i2] {
			ans[i] = s2[i2]
			i2++
		} else {
			ans[i] = s1[i1]
			i1++
		}
	}

	fmt.Print(string(ans))
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}
