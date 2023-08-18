/*
Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
https://en.wikipedia.org/wiki/Command_pattern

Тип: 		Поведенческий
Уровень: 	Объектный

Паттерн Command позволяет представить запрос в виде объекта. Из этого следует, что команда - это объект.
Такие запросы, например, можно ставить в очередь, отменять или возобновлять.

В этом паттерне мы оперируем следующими понятиями: Command - запрос в виде объекта на выполнение;
Receiver - объект-получатель запроса, который будет обрабатывать нашу команду; Invoker - объект-инициатор запроса.

Паттерн Command отделяет объект, инициирующий операцию, от объекта, который знает, как ее выполнить.
Единственное, что должен знать инициатор, это как отправить команду.

Требуется для реализации:
1. Базовый абстрактный класс Command описывающий интерфейс команды;
2. Класс ConcreteCommand, реализующий команду;
3. Класс Invoker, реализующий инициатора, записывающий команду и провоцирующий её выполнение;
4. Класс Receiver, реализующий получателя и имеющий набор действий, которые команда может запрашивать;

Invoker умеет складывать команды в стопку и инициировать их выполнение по какому-то событию.
Обратившись к Invoker можно отменить команду, пока та не выполнена.

ConcreteCommand содержит в себе запросы к Receiver, которые тот должен выполнять.

В свою очередь Receiver содержит только набор действий (Actions), которые выполняются при обращении к ним из ConcreteCommand.
*/
package main

// Предположим, у нас есть система умного дома, которая управляет различными устройствами, такими как свет, телевизор,
// кондиционери т.д. Мы хотим реализовать функционал управления устройствами с помощью паттерна "Команда".

// Сначала определим интерфейс "Комманда":
type Command interface {
	Execute()
	Undo()
}

// реализуем конкретные команды для различных действий:

// Light:

// Light On Command
type LightOnCommand struct {
	light *Light
}

func (c LightOnCommand) Execute() {
	c.light.On()
}

func (c LightOnCommand) Undo() {
	c.light.Off()
}

// Light off Command
type LightOffCommand struct {
	light *Light
}

func (c LightOffCommand) Execute() {
	c.light.Off()
}

func (c LightOffCommand) Undo() {
	c.light.On()
}

// TV:

// TV On Command
type TVOnCommand struct {
	tv *TV
}

func (c TVOnCommand) Execute() {
	c.tv.On()
}

func (c TVOnCommand) Undo() {
	c.tv.Off()
}

// TV Off Command
type TVOffCommand struct {
	tv *TV
}

func (c TVOffCommand) Execute() {
	c.tv.Off()
}

func (c TVOffCommand) Undo() {
	c.tv.On()
}

// реализуем объекты устройств:

// Light:
type Light struct {
	isOn bool
}

func (l Light) On() {
	l.isOn = true
}

func (l Light) Off() {
	l.isOn = false
}

// TV:
type TV struct {
	isOn bool
}

func (t TV) On() {
	t.isOn = true
}

func (t TV) Off() {
	t.isOn = false
}

// main
func main() {
	light := &Light{}
	lightOnCommand := &LightOnCommand{light: light}
	lightOffCommand := &LightOffCommand{light: light}

	tv := &TV{}
	tvOnCommand := &TVOnCommand{tv: tv}
	tvOffCommand := &TVOffCommand{tv: tv}

	// Выполняем команды
	lightOnCommand.Execute()  // Включить свет
	lightOffCommand.Execute() // Выключить свет

	tvOnCommand.Execute()  // Включить телевизор
	tvOffCommand.Execute() // Выключить телевизор

	// Отменяем последнюю выполненную команду
	lightOffCommand.Undo() // Вернуть состояние света на "включено"
	tvOffCommand.Undo()    // Вернуть состояние телевизора на "включено"
}
