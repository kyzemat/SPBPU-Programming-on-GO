package main

import (
	"fmt"
	"math"
)

// MaxIter — защита от бесконечного цикла при слишком малом ε.
const taylorMaxIter = 1000

// Sin — sin(x) рядом Тейлора: x - x³/3! + x⁵/5! - ...
func TaylorSin(x, eps float64) (float64, int) {
	x = math.Mod(x, 2*math.Pi) // приведение аргумента для точности

	term, sum, n := x, 0.0, 0
	for n < taylorMaxIter {
		sum += term
		n++
		term *= -x * x / float64((2*n)*(2*n+1))
		if math.Abs(term) < eps {
			break
		}
	}
	return sum, n
}

// Cos — cos(x) рядом Тейлора: 1 - x²/2! + x⁴/4! - ...
func TaylorCos(x, eps float64) (float64, int) {
	x = math.Mod(x, 2*math.Pi)

	term, sum, n := 1.0, 0.0, 0
	for n < taylorMaxIter {
		sum += term
		n++
		term *= -x * x / float64((2*n-1)*(2*n))
		if math.Abs(term) < eps {
			break
		}
	}
	return sum, n
}

// Exp — e^x рядом Тейлора: 1 + x + x²/2! + x³/3! + ...
func TaylorExp(x, eps float64) (float64, int) {
	term, sum, n := 1.0, 0.0, 0
	for n < taylorMaxIter {
		sum += term
		n++
		term *= x / float64(n)
		if math.Abs(term) < eps {
			break
		}
	}
	return sum, n
}

// Читаем аргумент x с консоли
func ReadX() float64 {
	for {
		var x float64

		fmt.Print("Введите x (например, 1.5): ")

		_, err := fmt.Scan(&x)
		if err != nil {
			fmt.Println("❌ Ошибка: введите число, например 1.5.")
			var trash string
			fmt.Scan(&trash)
			continue
		}

		return x
	}
}

// PrintTaylorTable — таблица для sin, cos, exp.
func PrintTaylorTable(x, eps float64) {
	type row struct {
		name  string
		f     func(float64, float64) (float64, int)
		exact float64
	}
	rows := []row{
		{"sin(x)", TaylorSin, math.Sin(x)},
		{"cos(x)", TaylorCos, math.Cos(x)},
		{"exp(x)", TaylorExp, math.Exp(x)},
	}

	fmt.Printf("\n─── Ряд Тейлора: x = %g, ε = %.0e ───\n", x, eps)
	fmt.Printf("%-8s │ %-16s │ %-16s │ %-12s │ %s\n", "Функция", "Приближённое", "Точное", "Погрешность", "Итераций")
	fmt.Println("─────────┼──────────────────┼──────────────────┼──────────────┼──────────")
	for _, r := range rows {
		approx, n := r.f(x, eps)
		fmt.Printf("%-8s │ %16.10f │ %16.10f │ %12.3e │ %d\n",
			r.name, approx, r.exact, math.Abs(approx-r.exact), n)
	}
}
