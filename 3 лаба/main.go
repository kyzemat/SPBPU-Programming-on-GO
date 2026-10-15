package main

import (
	"fmt"
	"math"
)

func main() {
	for {
		fmt.Println("\n========== МЕНЮ ==========")
		fmt.Println("1 — Ряд Лейбница для π")
		fmt.Println("2 — Ряд Тейлора для sin(x), cos(x), exp(x)")
		fmt.Println("0 — Выход")
		fmt.Print("Ваш выбор: ")

		var choice int
		if _, err := fmt.Scan(&choice); err != nil {
			fmt.Println("❌ Ошибка: введите 0, 1 или 2.")
			var trash string
			fmt.Scan(&trash)
			continue
		}

		switch choice {
		case 1:
			runLeibniz()
		case 2:
			runTaylor()
		case 0:
			fmt.Println("До свидания!")
			return
		default:
			fmt.Println("❌ Нет такого пункта меню. Введите 0, 1 или 2.")
		}
	}
}

// runTaylor — режим «Ряд Тейлора».
func runTaylor() {
	fmt.Println("\nРяд Тейлора для sin(x), cos(x), exp(x)")

	x := ReadX()
	eps := ReadEps()
	PrintTaylorTable(x, eps)
}

// runLeibniz — режим «Ряд Лейбница» (прежний код main).
func runLeibniz() {
	fmt.Println("\nРяд Лейбница: π = 4·(1 − 1/3 + 1/5 − 1/7 + ...)")

	// 1. Считываем точность с консоли
	eps := ReadEps()

	// 2. Считаем π с этой точностью
	PrintInfo(eps)

	// 3. Спрашиваем, нужен ли лимит итераций
	fmt.Print("\nОграничить число итераций? (y/n): ")
	var answer string
	fmt.Scan(&answer)

	if answer == "y" || answer == "Y" || answer == "д" || answer == "Д" {
		maxIter := ReadMaxIter()

		pi, iter, err, ok := LeibnizPiMaxIter(eps, maxIter)
		fmt.Println("\n========== С ЛИМИТОМ ИТЕРАЦИЙ ==========")
		fmt.Printf("Лимит:                  %d\n", maxIter)
		fmt.Printf("Приближённое значение:  π ≈ %.12f\n", pi)
		fmt.Printf("Точное значение:        π = %.12f\n", math.Pi)
		fmt.Printf("Погрешность:            %.3e\n", err)
		fmt.Printf("Итераций выполнено:     %d\n", iter)

		if ok {
			fmt.Println("✅ Точность достигнута.")
		} else {
			fmt.Println("⚠️  Лимит исчерпан, точность не достигнута.")
		}
		fmt.Println("========================================")
	}

	// 4. Таблица частичных сумм — как бонус
	fmt.Print("\nПоказать таблицу частичных сумм? (y/n): ")
	fmt.Scan(&answer)

	if answer == "y" || answer == "Y" || answer == "д" || answer == "Д" {
		PrintPartialSums([]int{1, 10, 100, 1000, 10000, 100000, 1000000})
	}
}
