package interfaces

type Shape interface {
	Area() int
}

type Square struct{
	side int
}

type Rectangle struct{
	length int
	breadth int
}

type Hybrid struct{
	//s Square
	//r Rectangle
	shapes []Shape
}

func (s Square) Area() int{
	return s.side*s.side
}

func (r Rectangle) Area() int{
	return r.length*r.breadth
}

func (h Hybrid) Area() int{
	//return h.s.Area()+h.r.Area()
	a:=0
	for _,i:=range h.shapes{
		a += i.Area()
	}
	return a
}