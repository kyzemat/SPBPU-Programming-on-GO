package shapes

// NewCircle — фабричная функция (конструктор) для Circle.
// Возвращает указатель, так как методы Circle имеют указательные ресиверы.
func NewCircle(radius float64) *Circle {
	return &Circle{Radius: radius}
}

// NewRectangle — фабричная функция (конструктор) для Rectangle.
func NewRectangle(width, height float64) *Rectangle {
	return &Rectangle{Width: width, Height: height}
}
