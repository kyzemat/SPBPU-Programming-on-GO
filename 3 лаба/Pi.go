package main

import (
	"fmt"
	"math"
)

const epsDefault = 1e-6

// LeibnizPi — вычисление π через ряд Лейбница до заданной точности.
func LeibnizPi(eps float64) (float64, int, float64) {
	if eps <= 0 {
		eps = epsDefault
	}

	sum := 0.0
	sign := 1.0
	denom := 1.0
	prev := 0.0
	iter := 0

	for {
		sum += sign / denom
		iter++

		pi := 4 * sum
		diff := math.Abs(pi - prev)

		if iter > 1 && diff < eps {
			return pi, iter, math.Abs(pi - math.Pi)
		}

		prev = pi
		sign = -sign
		denom += 2
	}
}

// LeibnizPiMaxIter — то же самое, но с ограничением на число итераций.
func LeibnizPiMaxIter(eps float64, maxIter int) (float64, int, float64, bool) {
	sum := 0.0
	sign := 1.0
	denom := 1.0
	iter := 0

	for iter < maxIter {
		sum += sign / denom
		iter++
		sign = -sign
		denom += 2

		pi := 4 * sum
		if math.Abs(pi-math.Pi) < eps {
			return pi, iter, math.Abs(pi - math.Pi), true
		}
	}

	pi := 4 * sum
	return pi, iter, math.Abs(pi - math.Pi), false
}

// LeibnizPartial — частичная сумма ровно n членов.
func LeibnizPartial(n int) float64 {
	sum := 0.0
	sign := 1.0
	denom := 1.0

	for i := 0; i < n; i++ {
		sum += sign / denom
		sign = -sign
		denom += 2
	}

	return 4 * sum
}

// PrintPartialSums — таблица частичных сумм.
func PrintPartialSums(limits []int) {
	fmt.Println("\n─── Частичные суммы ряда Лейбница ───")
	fmt.Printf("%10s │ %16s │ %14s │ %10s\n", "n членов", "π приближ.", "погрешность", "4·S(n)")
	fmt.Println("───────────┼──────────────────┼────────────────┼───────────")

	for _, n := range limits {
		pi := LeibnizPartial(n)
		fmt.Printf("%10d │ %16.10f │ %14.3e │ %10.4f\n",
			n, pi, math.Abs(pi-math.Pi), pi)
	}
}

// PrintInfo — общий отчёт по одному значению ε.
func PrintInfo(eps float64) {
	pi, iter, err := LeibnizPi(eps)

	fmt.Println("\n========== РЯД ЛЕЙБНИЦА ДЛЯ π ==========")
	fmt.Printf("Заданная точность:     ε = %.0e\n", eps)
	fmt.Printf("Приближённое значение: π ≈ %.12f\n", pi)
	fmt.Printf("Точное значение:       π = %.12f\n", math.Pi)
	fmt.Printf("Фактическая погрешность: %.3e\n", err)
	fmt.Printf("Число итераций:        %d\n", iter)
	fmt.Println("========================================")
}

// ReadMaxIter — читает максимальное число итераций.
func ReadMaxIter() int {
	for {
		var n int

		fmt.Print("Введите максимум итераций (например, 1000000): ")

		_, err := fmt.Scan(&n)
		if err != nil {
			fmt.Println("❌ Ошибка: введите целое число.")
			var trash string
			fmt.Scan(&trash)
			continue
		}

		if n <= 0 {
			fmt.Println("❌ Число итераций должно быть положительным.")
			continue
		}

		return n
	}
}
