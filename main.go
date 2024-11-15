package main

import (
	"fmt"
	"shapesmodule/figure"
)

func main() {
	// Создание различных фигур с помощью функции NewFigure
	square := figure.NewFigure("square", 4, 4)     // квадрат со стороной 4
	triangle := figure.NewFigure("triangle", 3, 5) // треугольник с основанием 3 и высотой 5
	circle := figure.NewFigure("circle", 7)        // круг с радиусом 7

	// Вызов функции FoundForm для каждой фигуры
	figure.FoundForm(square)
	figure.FoundForm(triangle)
	figure.FoundForm(circle)

	// Изменение стороны квадрата через функцию Change
	figure.Change(square.(*figure.Square), 10) // меняем сторону квадрата на 10

	// Выводим обновленную информацию о квадрате
	fmt.Println("После изменения стороны квадрата:")
	figure.FoundForm(square)
}
