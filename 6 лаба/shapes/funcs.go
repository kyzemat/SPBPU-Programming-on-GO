package shapes

import "fmt"

// PrintShapesInfo печатает информацию о каждой фигуре из списка.
// Функция работает с интерфейсом Shape и не знает конкретных типов:
func PrintShapesInfo(list []Shape) {
	for i, s := range list {
		fmt.Printf("%d. [%s] %s\n", i+1, s.Type(), s.String())
	}
}

// TotalArea возвращает суммарную площадь всех переданных фигур.
// Variadic-параметр позволяет вызывать её с любым числом фигур:
// TotalArea(circle, rect) или TotalArea(list...).
func TotalArea(list ...Shape) float64 {
	total := 0.0
	for _, s := range list {
		total += s.Area()
	}
	return total
}
