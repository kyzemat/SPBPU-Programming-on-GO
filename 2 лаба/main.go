package main

import "fmt"

func main() {
	for {
		fmt.Println("\n========== КАЛЬКУЛЯТОР ФИГУР ==========")
		fmt.Println("1 — Треугольник")
		fmt.Println("2 — Окружность")
		fmt.Println("0 — Выход")
		fmt.Print("Ваш выбор: ")

		var choice int
		if _, err := fmt.Scan(&choice); err != nil {
			fmt.Println("❌ Введите число.")
			var trash string
			fmt.Scan(&trash)
			continue
		}

		switch choice {
		case 1:
			t := CreateTriangle()
			t.PrintInfo()
		case 2:
			c := CreateCircle()
			c.PrintInfo()
			a := ReadAngle()
			c.PrintAngleInfo(a)
		case 0:
			fmt.Println("До свидания!")
			return
		default:
			fmt.Println("❌ Нет такого пункта.")
		}
	}
}
