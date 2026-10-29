package main

import (
	"OOP/shapes"
	"fmt"
)

func main() {
	fmt.Println("=== Лабораторная работа №6: Shape Interface ===")

	// Создаём фигуры через фабричные функции
	circle := shapes.NewCircle(5.0)
	rect := shapes.NewRectangle(4.0, 6.0)
	triangle := shapes.NewTriangle(3.0, 4.0, 5.0)

	// Демонстрация полиморфизма: все фигуры в одном срезе
	allShapes := []shapes.Shape{circle, rect, triangle}

	fmt.Println("\n--- Базовые фигуры ---")
	for _, s := range allShapes {
		fmt.Printf("[%s] %s\n", s.Type(), s.String())
	}

	// Демонстрация композиции через ColoredShape
	fmt.Println("\n--- Цветные фигуры (композиция через встраивание) ---")
	redCircle := shapes.NewColoredShape(circle, "Red")
	blueRect := shapes.NewColoredShape(rect, "Blue")
	greenTriangle := shapes.NewColoredShape(triangle, "Green")

	coloredShapes := []shapes.Shape{redCircle, blueRect, greenTriangle}
	for _, s := range coloredShapes {
		fmt.Printf("[%s] %s\n", s.Type(), s.String())
	}

	// Доказательство, что ColoredShape тоже реализует Shape
	fmt.Println("\n--- Проверка полиморфизма ---")
	var shape shapes.Shape = redCircle
	fmt.Printf("ColoredShape as Shape: Area = %.2f\n", shape.Area())
}
