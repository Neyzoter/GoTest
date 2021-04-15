package main

import "fmt"

type Point struct {
	x, y int
}

func main()  {
	p, q := Point{1, 2}, Point{3, 4}
	fmt.Println(p.minus(q))
}

func minus(p, q Point) Point {
	res := Point{0, 0}
	res.x = p.x - q.x
	res.y = p.y - q.y
	return res
}

func (p Point) minus(q Point) Point {
	res := Point{0, 0}
	res.x = p.x - q.x
	res.y = p.y - q.y
	return res
}
