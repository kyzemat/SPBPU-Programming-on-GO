package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"pricing_lab/pricing"
)

func main() {
	// Создаем "читатель" для ввода строк с клавиатуры
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n========== МЕНЮ: РАСЧЕТ ЦЕН ==========")
		fmt.Println("1 — Рассчитать цену с цепочкой модификаторов")
		fmt.Println("2 — Демонстрация замыканий (счетчик вызовов)")
		fmt.Println("3 — Демонстрация анонимных функций")
		fmt.Println("0 — Выход")
		fmt.Print("Ваш выбор: ")

		choiceStr, _ := reader.ReadString('\n')
		choiceStr = strings.TrimSpace(choiceStr)
		choice, err := strconv.Atoi(choiceStr)

		if err != nil {
			fmt.Println("❌ Ошибка: введите число (0, 1, 2 или 3).")
			continue
		}

		switch choice {
		case 1:
			runPriceCalculation(reader)
		case 2:
			runClosureDemo()
		case 3:
			runAnonymousDemo()
		case 0:
			fmt.Println("До свидания!")
			return
		default:
			fmt.Println("❌ Нет такого пункта меню. Введите 0, 1, 2 или 3.")
		}
	}
}

// runPriceCalculation — режим 1: интерактивный расчет цены с цепочкой модификаторов.
func runPriceCalculation(reader *bufio.Reader) {
	fmt.Println("\n--- Расчет цены ---")

	// 1. Считываем базовую цену
	fmt.Print("Введите базовую цену товара: ")
	basePrice := readFloat(reader)

	// 2. Считываем скидку
	fmt.Print("Введите процент скидки (0, если нет): ")
	discountPercent := readFloat(reader)

	// 3. Считываем налог
	fmt.Print("Введите процент налога (0, если нет): ")
	taxRate := readFloat(reader)

	// 4. Считываем точность округления
	fmt.Print("Введите точность округления (кол-во знаков после запятой): ")
	precision := readInt(reader)

	// 5. Считываем минимальную цену
	fmt.Print("Введите минимальную цену (0, если не нужна): ")
	minPrice := readFloat(reader)

	// Собираем модификаторы в слайс
	var modifiers []pricing.PriceModifier

	// Добавляем только те модификаторы, которые имеют смысл
	if discountPercent > 0 {
		modifiers = append(modifiers, pricing.WithDiscount(discountPercent))
	}
	if taxRate > 0 {
		modifiers = append(modifiers, pricing.WithTax(taxRate))
	}
	if precision >= 0 {
		modifiers = append(modifiers, pricing.WithRounding(precision))
	}
	if minPrice > 0 {
		modifiers = append(modifiers, pricing.WithMinimumPrice(minPrice))
	}

	// Если список пуст — возвращаем исходную цену
	chain := pricing.CombineModifiers(modifiers...)
	finalPrice := chain(basePrice)

	// Выводим результат
	fmt.Println("\n========== РЕЗУЛЬТАТ ==========")
	fmt.Printf("Базовая цена:       %.2f\n", basePrice)
	fmt.Printf("Скидка:             %.1f%%\n", discountPercent)
	fmt.Printf("Налог:              %.1f%%\n", taxRate)
	fmt.Printf("Точность округл.:   %d знаков\n", precision)
	fmt.Printf("Мин. цена:          %.2f\n", minPrice)
	fmt.Println("-------------------------------")
	fmt.Printf("ИТОГОВАЯ ЦЕНА:      %.2f\n", finalPrice)
	fmt.Println("================================")
}

// runClosureDemo — режим 2: демонстрация замыканий (счетчик вызовов).
func runClosureDemo() {
	fmt.Println("\n--- Демонстрация замыканий: счетчик вызовов ---")

	counter := 0
	incrementPrice := func(p float64) float64 {
		counter++
		fmt.Printf("  [Модификатор вызван %d раз(а)]\n", counter)
		return p + 10
	}

	chain := pricing.CombineModifiers(incrementPrice, incrementPrice)
	result := chain(100.0)

	fmt.Printf("\nРезультат с счетчиком: %.2f\n", result)
	fmt.Printf("Итоговое число вызовов: %d\n", counter)
	fmt.Println("Как видите, переменная counter захвачена замыканием и изменяется между вызовами.")
}

// runAnonymousDemo — режим 3: демонстрация анонимных функций.
func runAnonymousDemo() {
	fmt.Println("\n--- Демонстрация анонимных функций ---")

	fmt.Print("Введите цену: ")
	basePrice := readFloat(bufio.NewReader(os.Stdin))

	// Анонимная функция прямо в аргументе CombineModifiers
	modifier := pricing.CombineModifiers(
		func(p float64) float64 {
			return p + 50 // наценка 50
		},
		pricing.WithRounding(0), // округление до целого
	)

	finalPrice := modifier(basePrice)
	fmt.Printf("\nЦена %.2f с наценкой 50 и округлением: %.0f\n", basePrice, finalPrice)
	fmt.Println("Анонимная функция создана прямо в месте вызова, без имени.")
}

// --- Вспомогательные функции для безопасного ввода ---

// readFloat считывает число с плавающей точкой с защитой от ошибок.
func readFloat(reader *bufio.Reader) float64 {
	for {
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		val, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Print("❌ Ошибка. Введите число: ")
			continue
		}
		return val
	}
}

// readInt считывает целое число с защитой от ошибок.
func readInt(reader *bufio.Reader) int {
	for {
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		val, err := strconv.Atoi(input)
		if err != nil {
			fmt.Print("❌ Ошибка. Введите целое число: ")
			continue
		}
		return val
	}
}
