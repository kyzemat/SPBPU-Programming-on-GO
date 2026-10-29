package shapes

import "fmt"

// ColoredShape — декоратор, добавляющий цвет к любой фигуре.
// Демонстрирует композицию в Go: структура встраивает интерфейс Shape
// и получает все его методы автоматически, при этом добавляя свои поля.
type ColoredShape struct {
	Shape        // Встраивание интерфейса: ColoredShape "наследует" Area(), Perimeter(), Type()
	Color string // Дополнительное поле
}

// String переопределяет метод String() для добавления информации о цвете.
// Обратите внимание: вызов c.Shape.String() обращается к методу исходной фигуры.
func (cs *ColoredShape) String() string {
	return fmt.Sprintf("[%s] %s", cs.Color, cs.Shape.String())
}

// NewColoredShape — фабричная функция (конструктор) для ColoredShape.
// Возвращает указатель на структуру для единообразия с остальными конструкторами.
func NewColoredShape(s Shape, color string) *ColoredShape {
	return &ColoredShape{
		Shape: s,
		Color: color,
	}
}
