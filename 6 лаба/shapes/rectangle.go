package shapes

import "fmt"

// Rectangle описывает прямоугольник.
// Используем значимые ресиверы (value receivers), потому что:
// 1. Структура маленькая (2 поля float64).
// 2. Методы не изменяют состояние объекта.
// 3. Копирование дешевле, чем разыменование указателя.
type Rectangle struct {
	Width, Height float64
}

// Area возвращает площадь прямоугольника.
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Perimeter возвращает периметр прямоугольника.
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// String возвращает строковое представление прямоугольника.
func (r Rectangle) String() string {
	return fmt.Sprintf("Rectangle(Width=%.2f, Height=%.2f, Area=%.2f, Perimeter=%.2f)",
		r.Width, r.Height, r.Area(), r.Perimeter())
}

// Type возвращает тип фигуры.
func (r Rectangle) Type() string {
	return "Rectangle"
}
