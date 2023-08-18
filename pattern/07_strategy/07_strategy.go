/*
Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
https://en.wikipedia.org/wiki/Strategy_pattern

Тип: 		Поведенческий
Уровень: 	Объектный

Паттерн Strategy определяет набор алгоритмов схожих по роду деятельности, инкапсулирует их в отдельный класс и делает
их подменяемыми. Паттерн Strategy позволяет подменять алгоритмы без участия клиентов, которые используют эти алгоритмы.

Требуется для реализации:
Абстрактный класс Strategy, определяющий интерфейс различных стратегий;
Класс StrategyA, реализует одну из стратегий представляющую собой алгоритмы, направленные на достижение определенной цели;
Класс StrategyB, реализует одно из стратегий представляющую собой алгоритмы, направленные на достижение определенной цели.

Применимость паттерна Стратегия:
- Когда есть несколько взаимозаменяемых вариантов решения задачи и необходимо выбрать один из них.
- Когда необходимо скрыть детали реализации алгоритма от клиентского кода.
- Когда есть необходимость добавления новых алгоритмов или изменения существующих без изменения клиентского кода.

Плюсы использования паттерна Стратегия:
  - Улучшение поддерживаемости: Каждый алгоритм инкапсулируется в отдельном классе, что упрощает его понимание и изменение.
  - Гибкость и расширяемость: Позволяет добавлять новые алгоритмы без изменения существующего кода,
    а также легко комбинировать и переключаться между ними.
  - Уменьшение связанности: Клиентский код зависит только от абстрактного класса или интерфейса стратегии,
    что уменьшает связанность между компонентами системы.

Минусы использования паттерна Стратегия:
  - Усложнение структуры программы: Для каждого алгоритма требуется создание отдельного класса стратегии,
    что может привести к увеличению количества классов и усложнению кода.
  - Дополнительные затраты на память: Каждый класс стратегии требует дополнительной памяти для хранения своего состояния.

Реальные примеры использования паттерна Стратегия:
- Сортировка данных
- Операции с файлами
- Парсинг данных
*/
package main

import "fmt"

// Абстрактный класс Strategy, определяющий интерфейс различных стратегий;
type Payment interface {
	Pay() error
}

// Класс StrategyA, реализует одну из стратегий
type cardPayment struct {
	cardNumber string
	cvv        string
}

func NewCardPayment(cardNumber, cvv string) Payment {
	return &cardPayment{
		cardNumber: cardNumber,
		cvv:        cvv,
	}
}

func (p *cardPayment) Pay() error {
	// ... implementation
	return nil
}

// Класс StrategyB, реализует одну из стратегий
type payPalPayment struct {
}

func NewPayPalPayment() Payment {
	return &payPalPayment{}
}

func (p *payPalPayment) Pay() error {
	// ... implementation
	return nil
}

// Класс StrategyC, реализует одну из стратегий
type qiwiPayment struct {
}

func NewQIWIPayment() Payment {
	return &qiwiPayment{}
}

func (p *qiwiPayment) Pay() error {
	// ... implementation
	return nil
}

// ------------------------------
func Strategies(product string, payWay int) {
	var payment Payment
	switch payWay {
	case 1:
		payment = NewCardPayment("1234 5678 9101 1121", "123")
	case 2:
		payment = NewPayPalPayment()
	case 3:
		payment = NewQIWIPayment()
	}

	processOrder(product, payment)
}

func processOrder(product string, payment Payment) {
	// ... implementation
	err := payment.Pay()
	if err != nil {
		return
	}
	fmt.Println(product, payment)
}

// Main
func main() {
	Strategies("ticket", 1)
	Strategies("bred", 2)
	Strategies("car", 3)
}
