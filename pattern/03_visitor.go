/*
Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
https://en.wikipedia.org/wiki/Visitor_pattern

Тип: Поведенческий
Уровнь: Объектный

Паттерн Visitor позволяет обойти набор элементов (объектов) с разнородными интерфейсами,
а также позволяет добавить новый метод в класс объекта, при этом, не изменяя сам класс этого объекта.

Требуется для реализации:
1. Абстрактный класс Visitor, описывающий интерфейс визитера;
2. Класс ConcreteVisitor, реализующий конкретного визитера. Реализует методы для обхода конкретного элемента;
3. Класс ObjectStructure, реализующий структуру(коллекцию), в которой хранятся элементы для обхода;
4. Абстрактный класс Element, реализующий интерфейс элементов структуры;
5. Класс ElementA, реализующий элемент структуры;
6. Класс ElementB, реализующий элемент структуры.


*/

package pattern

import "fmt"

// Абстрактный класс Visitor, описывающий интерфейс визитера;
type Visitor interface {
	VisitBar()
	VisitCinema()
	VisitEmbassy()
}

// Класс ConcreteVisitor, реализующий конкретного визитера. Реализует методы для обхода конкретного элемента;
type ConcreteVisitor struct{}

func (cv *ConcreteVisitor) VisitBar() {
	fmt.Println("Visit Bar!")
}

func (cv *ConcreteVisitor) VisitCinema() {
	fmt.Println("Visit Cinema!")
}

func (cv *ConcreteVisitor) VisitEmbassy() {
	fmt.Println("Visit Embassy!")
}

// Класс ObjectStructure, реализующий структуру(коллекцию), в которой хранятся элементы для обхода;
type City struct {
	places []Place
}

// Абстрактный класс Element, реализующий интерфейс элементов структуры;
type Place interface {
	Accept(v Visitor) string
}
