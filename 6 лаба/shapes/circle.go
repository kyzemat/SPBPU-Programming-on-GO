package shapes

import (
	"fmt"
	"math"
)

// Circle описывает круг.
// Используем указательные ресиверы (pointer receivers), потому что:
// 1. Структура содержит поля, которые логически принадлежат объекту.
// 2. В будущем может быть расширена (например, кэширование площади).
// 3. Это позволяет избежать копирования структуры при каждом вызове.
type Circle struct {
	Radius float64
}

// Area возвращает площадь круга: π * r².
func (c *Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Perimeter возвращает длину окружности: 2 * π * r.
func (c *Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// String возвращает строковое представление круга.
func (c *Circle) String() string {
	return fmt.Sprintf("Circle(Radius=%.2f, Area=%.2f, Perimeter=%.2f)",
		c.Radius, c.Area(), c.Perimeter())
}

// Type возвращает тип фигуры.
func (c *Circle) Type() string {
	return "Circle"
}
