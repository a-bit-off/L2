/*
Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
https://en.wikipedia.org/wiki/State_pattern

Тип:		Поведенческий
Уровень:	Объектный

Паттерн "Состояние" применяется, когда у объекта есть внутреннее состояние, которое влияет на его поведение, и поведение
объекта должно изменяться в зависимости от этого состояния. Он позволяет объекту динамически менять свое поведение без
необходимости изменения его класса.

Требуется для реализации:
1. Класс Context, представляет собой объектно-ориентированное представление конечного автомата;
2. Абстрактный класс State, определяющий интерфейс различных состояний;
3. Класс ConcreteStateA реализует одно из поведений, ассоциированное с определенным состоянием;
4. Класс ConcreteStateB реализует одно из поведений, ассоциированное с определенным состоянием.

Паттерн должен применяться:
- когда поведение объекта зависит от его состояния
- поведение объекта должно изменяться во время выполнения программы
- состояний достаточно много и использовать для этого условные операторы, разбросанные по коду, достаточно затруднительно

Преимущества паттерна "Состояние":
 1. Упрощает код объекта, так как каждое состояние имеет свой класс и отвечает только за свою функциональность.
 2. Объект и его состояния могут изменяться независимо друг от друга, что облегчает добавление новых состояний или
    изменение существующих.
 3. Избавляет от больших условных операторов, которые проверяют текущее состояние объекта.

Недостатки паттерна "Состояние":
1. Может привести к увеличению числа классов в системе, особенно если состояний много.
2. Может быть сложно следить за переходами между состояниями и поддерживать их целостность.

Реальные примеры использования паттерна "Состояние":
1. Автомат продажи билетов: состояниями могут быть "ожидание ввода денег", "выбор места", "оплата", "печать билета".
2. Редактор текста: состояниями могут быть "режим вставки", "режим выделения", "режим редактирования".
3. Игра с различными уровнями сложности: состояниями могут быть "легкий уровень", "средний уровень", "сложный уровень".
*/
package main

import "fmt"

// рассмотрим реализацию паттерна "Состояние" на языке Go на примере автомата продажи напитков.

// Абстрактный класс State, определяющий интерфейс различных состояний;
type State interface {
	InsertCoin()    // Вставьте монету
	EjectCoin()     // Извлечь монету
	SelectDrink()   // Выбрать напиток
	DispenseDrink() // Дозировать напиток
}

// Классы реализуют одно из поведений, ассоциированное с определенным состоянием;

// No Coin State
type NoCoinState struct{}

func (s *NoCoinState) InsertCoin() {
	fmt.Println("Монетка вставлена")
}

func (s *NoCoinState) EjectCoin() {
	fmt.Println("Монетка не вставлена")
}

func (s *NoCoinState) SelectDrink() {
	fmt.Println("Вставьте монетку")
}

func (s *NoCoinState) DispenseDrink() {
	fmt.Println("Вставьте монетку")
}

// Coin Inserted State
type CoinInsertedState struct{}

func (s *CoinInsertedState) InsertCoin() {
	fmt.Println("Монетка уже вставлена")
}

func (s *CoinInsertedState) EjectCoin() {
	fmt.Println("Монетка извлечена")
}

func (s *CoinInsertedState) SelectDrink() {
	fmt.Println("Выберете напиток")
}

func (s *CoinInsertedState) DispenseDrink() {
	fmt.Println("Вставьте монетку")
}

// Drink Selected State
type DrinkSelectedState struct{}

func (s *DrinkSelectedState) InsertCoin() {
	fmt.Println("Монетка уже вставлена")
}

func (s *DrinkSelectedState) EjectCoin() {
	fmt.Println("Монетка извлечена")
}

func (s *DrinkSelectedState) SelectDrink() {
	fmt.Println("Напиток уже выбран")
}

func (s *DrinkSelectedState) DispenseDrink() {
	fmt.Println("Напиток выдается")
}

// Класс Context, представляет собой объектно-ориентированное представление конечного автомата;

type VendingMachine struct {
	state State
}

func NewVendingMachine() *VendingMachine {
	return &VendingMachine{state: &NoCoinState{}}
}

func (m *VendingMachine) setState(state State) {
	m.state = state
}

func (m *VendingMachine) InsertCoin() {
	m.state.InsertCoin()
	m.setState(&CoinInsertedState{})
}

func (m *VendingMachine) EjectCoin() {
	m.state.EjectCoin()
	m.setState(&NoCoinState{})
}

func (m *VendingMachine) SelectDrink() {
	m.state.SelectDrink()
	m.setState(&DrinkSelectedState{})
}

func (m *VendingMachine) DispenseDrink() {
	m.state.DispenseDrink()
	m.setState(&NoCoinState{})
}

func main() {
	vendingMachine := NewVendingMachine()
	vendingMachine.InsertCoin()    // Монетка вставлена
	vendingMachine.SelectDrink()   // Вставьте монетку
	vendingMachine.EjectCoin()     // Монетка извлечена
	vendingMachine.InsertCoin()    // Монетка вставлена
	vendingMachine.SelectDrink()   // Выберите напиток
	vendingMachine.DispenseDrink() // Напиток выдается
	vendingMachine.SelectDrink()   // Напиток уже выбран
	vendingMachine.EjectCoin()     // Монетка извлечена
}
