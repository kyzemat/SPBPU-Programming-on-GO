package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Ряд Лейбница: π = 4·(1 − 1/3 + 1/5 − 1/7 + ...)")

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
