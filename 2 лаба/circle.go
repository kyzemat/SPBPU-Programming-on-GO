package main

import (
	"fmt"
	"math"
)

type Circle struct {
	R float64
}

// Проверка существования окружности
func (c Circle) IsValid() bool {
	return c.R > 0
}

// Диаметр
func (c Circle) Diameter() float64 {
	return 2 * c.R
}

// Длина окружности
func (c Circle) Circumference() float64 {
	return 2 * math.Pi * c.R
}

// Площадь круга
func (c Circle) Area() float64 {
	return math.Pi * c.R * c.R
}

// Длина дуги для угла в градусах
func (c Circle) ArcLength(angleDeg float64) float64 {
	return math.Pi * c.R * angleDeg / 180
}

// Площадь сектора для угла в градусах
func (c Circle) SectorArea(angleDeg float64) float64 {
	return math.Pi * c.R * c.R * angleDeg / 360
}

// Длина хорды, стягивающей дугу
func (c Circle) ChordLength(angleDeg float64) float64 {
	return 2 * c.R * math.Sin(angleDeg*math.Pi/360)
}

// Классификация угла — демонстрация switch
func ClassifyAngle(angleDeg float64) string {
	switch {
	case angleDeg < 0 || angleDeg > 360:
		return "некорректный"

	case equal(angleDeg, 0):
		return "нулевой"

	case angleDeg < 90:
		return "острый"

	case equal(angleDeg, 90):
		return "прямой"

	case angleDeg < 180:
		return "тупой"

	case equal(angleDeg, 180):
		return "развёрнутый"

	case angleDeg < 360:
		return "рефлекторный"

	default:
		return "полный"
	}
}

// Вывод информации об окружности
func (c Circle) PrintInfo() {
	fmt.Println("\n========== ОКРУЖНОСТЬ ==========")

	fmt.Printf("Радиус: %.2f\n", c.R)

	if !c.IsValid() {
		fmt.Println("❌ Такой окружности не существует :(")
		fmt.Println("Радиус должен быть больше нуля")
		return
	}

	fmt.Println("✅ Окружность существует.")

	fmt.Printf("Диаметр: %.2f\n", c.Diameter())
	fmt.Printf("Длина окружности: %.2f\n", c.Circumference())
	fmt.Printf("Площадь круга: %.2f\n", c.Area())

	fmt.Println("=================================")
}

// Вывод информации о дуге/секторе/хорде
func (c Circle) PrintAngleInfo(angleDeg float64) {
	fmt.Println("\n========== УГОЛ И ДУГА ==========")

	if angleDeg < 0 || angleDeg > 360 {
		fmt.Println("❌ Угол должен быть в диапазоне от 0 до 360 градусов.")
		return
	}

	fmt.Printf("Угол: %.2f°\n", angleDeg)
	fmt.Printf("Тип угла: %s\n", ClassifyAngle(angleDeg))
	fmt.Printf("Длина дуги: %.2f\n", c.ArcLength(angleDeg))
	fmt.Printf("Площадь сектора: %.2f\n", c.SectorArea(angleDeg))
	fmt.Printf("Длина хорды: %.2f\n", c.ChordLength(angleDeg))

	fmt.Println("=================================")
}

// Чтение положительного числа с клавиатуры
func ReadPositive(name string) float64 {
	for {
		var value float64

		fmt.Printf("Введите %s: ", name)

		_, err := fmt.Scan(&value)

		if err != nil {
			fmt.Println("❌ Ошибка: введите число.")

			var trash string
			fmt.Scan(&trash)

			continue
		}

		if value <= 0 {
			fmt.Println("❌ Значение должно быть больше нуля.")
			continue
		}

		return value
	}
}

// Чтение угла в диапазоне [0; 360]
func ReadAngle() float64 {
	for {
		var value float64

		fmt.Print("Введите угол в градусах (0..360): ")

		_, err := fmt.Scan(&value)

		if err != nil {
			fmt.Println("❌ Ошибка: введите число.")

			var trash string
			fmt.Scan(&trash)

			continue
		}

		if value < 0 || value > 360 {
			fmt.Println("❌ Угол должен быть в диапазоне от 0 до 360.")
			continue
		}

		return value
	}
}

// Создание окружности
func CreateCircle() Circle {
	for {
		r := ReadPositive("радиус")

		circle := Circle{
			R: r,
		}

		if !circle.IsValid() {
			fmt.Println("\n❌ Такая окружность не существует.")
			fmt.Println("Радиус должен быть больше нуля.")
			fmt.Println("Попробуйте ввести радиус ещё раз.\n")
			continue
		}

		return circle
	}
}
