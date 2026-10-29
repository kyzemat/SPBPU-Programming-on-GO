package shapes

import (
	"fmt"
	"math"
)

// Triangle описывает треугольник по трём сторонам.
// Используем значимые ресиверы (value receivers), потому что структура
// маленькая (3 поля float64), а методы не изменяют состояние объекта.
type Triangle struct {
	A, B, C float64
}

// NewTriangle — фабричная функция (конструктор) для Triangle.
// Стороны должны быть положительными и удовлетворять неравенству
// треугольника (сумма любых двух сторон больше третьей), иначе
// функция вызывает panic: это ошибка программиста, а не пользователя.
func NewTriangle(a, b, c float64) *Triangle {
	if a <= 0 || b <= 0 || c <= 0 || a+b <= c || a+c <= b || b+c <= a {
		panic(fmt.Sprintf("shapes: треугольник со сторонами %.2f, %.2f, %.2f не существует", a, b, c))
	}
	return &Triangle{A: a, B: b, C: c}
}

// Area возвращает площадь по формуле Герона.
func (t Triangle) Area() float64 {
	p := t.Perimeter() / 2
	return math.Sqrt(p * (p - t.A) * (p - t.B) * (p - t.C))
}

// Perimeter возвращает периметр треугольника.
func (t Triangle) Perimeter() float64 {
	return t.A + t.B + t.C
}

// String возвращает строковое представление треугольника.
func (t Triangle) String() string {
	return fmt.Sprintf("Triangle(A=%.2f, B=%.2f, C=%.2f, Area=%.2f, Perimeter=%.2f)",
		t.A, t.B, t.C, t.Area(), t.Perimeter())
}

// Type возвращает тип фигуры.
func (t Triangle) Type() string {
	return "Triangle"
}
