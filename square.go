package square

type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func (m Square) End() Point {
	// implement me
	var end Point
	end.x = m.start.x + int(m.a)
	end.y = m.start.y + int(m.a)

	return end
}

func (m Square) Area() uint {
	// implement me
	return m.a * m.a
}

func (m Square) Perimeter() uint {
	// implement me
	return 4 * m.a
}
