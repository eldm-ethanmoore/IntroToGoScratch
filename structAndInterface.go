package main

import ("fmt"; "math")

type MultiShape struct {
	shapes []Shape
}

type Shape interface {
	area() float64
}

type Person struct {
	Name string
}

type Android struct {
	Person
	Model string
}

type Circle struct {
	x, y, r float64
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}

func (p *Person) Talk() {
	fmt.Println("My name is ", p.Name)
}

func (c *Circle) area() float64 {
	return math.Pi * c.r*c.r
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

func (m *MultiShape) area() float64 {
	var area float64
	for _, s := range m.shapes {
		area += s.area()
	}
	return area
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

func main() {
	//c := new(Circle)
	//c := Circle{x: 0, y: 0, r: 5}
	c := Circle{0, 0, 5}
	//fmt.Println(c.y, c.r)
	fmt.Println(c.area())
	r := Rectangle{1, 1, 20, 20}
	fmt.Println(r.area())
	a := new(Android)
	a.Name = "eldm"
	a.Talk()
	fmt.Println(totalArea(&c, &r))
}