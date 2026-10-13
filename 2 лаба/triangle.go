package main

import (
	"fmt"
	"math"
)

const eps = 1e-9

type Triangle struct {
	A float64
	B float64
	C float64
}

// Проверка существования треугольника
func (t Triangle) IsValid() bool {
	return t.A > 0 &&
		t.B > 0 &&
		t.C > 0 &&
		t.A+t.B > t.C &&
		t.A+t.C > t.B &&
		t.B+t.C > t.A
}

// Периметр
func (t Triangle) Perimeter() float64 {
	return t.A + t.B + t.C
}

// Площадь по формуле Герона
func (t Triangle) Area() float64 {
	p := t.Perimeter() / 2

	return math.Sqrt(
		p * (p - t.A) *
			(p - t.B) *
			(p - t.C),
	)
}

// Сравнение чисел с учётом погрешности float64
func equal(a, b float64) bool {
	return math.Abs(a-b) < eps
}

// Определение типа по сторонам
func (t Triangle) TypeBySides() string {
	switch {
	case equal(t.A, t.B) && equal(t.B, t.C):
		return "равносторонний"

	case equal(t.A, t.B) ||
		equal(t.A, t.C) ||
		equal(t.B, t.C):
		return "равнобедренный"

	default:
		return "разносторонний"
	}
}

// Определение типа треугольника по углам
func (t Triangle) TypeByAngles() string {
	a, b, c := t.A, t.B, t.C

	// Сротировка сторон по возрастанию
	if a > b {
		a, b = b, a
	}

	if b > c {
		b, c = c, b
	}

	if a > b {
		a, b = b, a
	}

	left := a*a + b*b
	right := c * c

	switch {
	case math.Abs(left-right) < eps:
		return "прямоугольный"

	case left > right:
		return "остроугольный"

	default:
		return "тупоугольный"
	}
}

// Вывод информации о треугольнике
func (t Triangle) PrintInfo() {
	fmt.Println("\n========== ТРЕУГОЛЬНИК ==========")

	fmt.Printf("Сторона A: %.2f\n", t.A)
	fmt.Printf("Сторона B: %.2f\n", t.B)
	fmt.Printf("Сторона C: %.2f\n", t.C)

	if !t.IsValid() {
		fmt.Println("❌ Такого треугольника не существует :(")
		fmt.Println("Сумма любых двух сторон должна быть больше третьей")
		return
	}

	fmt.Println("✅ Треугольник существует.")

	fmt.Printf("Периметр: %.2f\n", t.Perimeter())
	fmt.Printf("Площадь: %.2f\n", t.Area())
	fmt.Printf("По сторонам это: %s\n", t.TypeBySides())
	fmt.Printf("По углам это: %s\n", t.TypeByAngles())

	fmt.Println("=================================")
}

func ReadSide(name string) float64 {
	for {
		var side float64

		fmt.Printf("Введите сторону %s: ", name)

		_, err := fmt.Scan(&side)

		if err != nil {
			fmt.Println("❌ Ошибка: введите число.")

			var trash string
			fmt.Scan(&trash)

			continue
		}

		if side <= 0 {
			fmt.Println("❌ Сторона должна быть больше нуля.")
			continue
		}

		return side
	}
}

// Создание треугольника
func CreateTriangle() Triangle {
	for {
		a := ReadSide("A")
		b := ReadSide("B")
		c := ReadSide("C")

		triangle := Triangle{
			A: a,
			B: b,
			C: c,
		}

		if !triangle.IsValid() {
			fmt.Println("\n❌ Такой треугольник не существует.")
			fmt.Println("Сумма любых двух сторон должна быть больше третьей.")
			fmt.Println("Попробуйте ввести стороны ещё раз.\n")
			continue
		}

		return triangle
	}
}

func main() {

	triangle := CreateTriangle()
	triangle.PrintInfo()
}
