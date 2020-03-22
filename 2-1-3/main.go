package main

import "fmt"

var n = 10
var m = 10
var field = [][]string {
	{"#", "S", "#", "#", "#", "#", "#", "#", ".", "#"},
	{".", ".", ".", ".", ".", ".", "#", ".", ".", "#"},
	{".", "#", ".", "#", "#", ".", "#", "#", ".", "#"},
	{".", "#", ".", ".", ".", ".", ".", ".", ".", "."},
	{"#", "#", ".", "#", "#", ".", "#", "#", "#", "#"},
	{".", ".", ".", ".", "#", ".", ".", ".", ".", "#"},
	{".", "#", "#", "#", "#", "#", "#", "#", ".", "#"},
	{".", ".", ".", ".", "#", ".", ".", ".", ".", "."},
	{".", "#", "#", "#", "#", ".", "#", "#", "#", "."},
	{".", ".", ".", ".", "#", ".", ".", ".", "G", "#"},
}

const INF = 100000
var d [][]int
var sx, sy int
var gx, gy int

type position struct {
	x int
	y int
}

type queue []position
var q queue

var vx = []int{1, 0, -1, 0}
var vy = []int{0, 1, 0, -1}

func (q *queue) push(v position) {
	*q = append(*q, v)
}

func (q *queue) pop() position {
	_q := *q
	res := _q[0]
	*q = _q[1:]
	return res
}

func main() {
	q = queue{}

	d = make([][]int, m)
	for i := range d {
		 d[i] = make([]int, n)
	}

	for i:=0; i<n; i++ {
		for j:=0; j<m; j++ {
			d[i][j] = INF
			if field[i][j] == "S" {
				sx, sy = i, j
			}
			if field[i][j] == "G" {
				gx, gy = i, j
			}
		}
	}

	q.push(position{sx, sy})
	d[sx][sy] = 0

	for len(q) >= 0 {
		p := q.pop()
		if p.x == gx && p.y == gy {
			break
		}

		for i:=0; i<4; i++ {
			nx, ny := p.x + vx[i], p.y + vy[i]

			if nx >= 0 && nx < n && ny >= 0 && ny < m && field[nx][ny] != "#" && d[nx][ny] == INF {
				d[nx][ny] = d[p.x][p.y] + 1
				q.push(position{nx, ny})
			}

		}
	}
	fmt.Print(d[gx][gy])
}
