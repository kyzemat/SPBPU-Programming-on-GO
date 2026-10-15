package main

import (
	"fmt"
)

// ReadEps — читает точность ε с консоли с валидацией.
// Допускает ввод в формате 1e-6, 0.000001, 1E-3.
func ReadEps() float64 {
	for {
		var eps float64

		fmt.Print("Введите точность ε (например, 1e-6): ")

		_, err := fmt.Scan(&eps)
		if err != nil {
			fmt.Println("❌ Ошибка: введите число в формате 1e-6 или 0.000001.")

			// очистка буфера ввода
			var trash string
			fmt.Scan(&trash)

			continue
		}

		if eps <= 0 {
			fmt.Println("❌ Точность должна быть больше нуля.")
			continue
		}

		if eps >= 1 {
			fmt.Println("❌ Точность должна быть меньше 1 (например, 1e-3).")
			continue
		}

		return eps
	}
}
