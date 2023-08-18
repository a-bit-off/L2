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

Плюсы использования паттерна "Команда":
- Разделение отправителя и получателя: Команда инкапсулирует запрос на выполнение операции и отделяет отправителя запроса от получателя.
- Поддержка отмены операций: Команды могут иметь методы отмены, что позволяет отменять выполненные операции.
- Поддержка истории выполненных операций: Команды могут сохраняться в истории, что позволяет восстанавливать предыдущие состояния системы.

Минусы использования паттерна "Команда":
- Усложнение кода: Внедрение команд может привести к увеличению количества классов и усложнению структуры кода.
- Дополнительные затраты на память: Использование истории выполненных команд может потребовать дополнительной памяти для их хранения.

Реальные примеры использования паттерна "Команда" на практике:
- Редактор текста: Команды могут использоваться для выполнения операций редактирования текста, таких как вставка, удаление, отмена/повтор и т.д.
- Умный дом: Команды могут использоваться для управления различными устройствами в умном доме, такими как включение/выключение света, регулировка температуры и т.д.
- Интерфейс пользователя: Команды могут использоваться для выполнения операций, связанных с действиями пользователя, такими как нажатие кнопок, выбор пунктов меню и т.д.
*/
package main

import "fmt"

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

// TV Turn Up Volume Command
type TVTurnUpVolumeCommand struct {
	tv *TV
}

func (c TVTurnUpVolumeCommand) Execute() {
	c.tv.UpVolume()
}

func (c TVTurnUpVolumeCommand) Undo() {
	c.tv.DownVolume()
}

// TV Turn Down Volume Command
type TVTurnDownVolumeCommand struct {
	tv *TV
}

func (c TVTurnDownVolumeCommand) Execute() {
	c.tv.DownVolume()
}

func (c TVTurnDownVolumeCommand) Undo() {
	c.tv.UpVolume()
}

// реализуем объекты устройств:

// Light:
type Light struct {
	isOn bool
}

func (l *Light) On() {
	l.isOn = true
}

func (l *Light) Off() {
	l.isOn = false
}

// TV:
type TV struct {
	isOn   bool
	volume int
}

func (t *TV) On() {
	t.isOn = true
}

func (t *TV) Off() {
	t.isOn = false
}

func (t *TV) UpVolume() {
	t.volume++
}

func (t *TV) DownVolume() {
	t.volume--
}

// main
func main() {
	light := &Light{}
	lightOnCommand := &LightOnCommand{light: light}
	lightOffCommand := &LightOffCommand{light: light}

	tv := &TV{}
	tvOnCommand := &TVOnCommand{tv: tv}
	tvOffCommand := &TVOffCommand{tv: tv}

	tvTurnUpVolumeCommand := TVTurnUpVolumeCommand{tv: tv}
	tvTurnDownVolumeCommand := TVTurnDownVolumeCommand{tv: tv}

	// Выполняем команды
	lightOnCommand.Execute() // Включить свет
	fmt.Println("light:", light)

	lightOffCommand.Execute() // Выключить свет
	fmt.Println("light:", light)

	tvOnCommand.Execute() // Включить телевизор
	fmt.Println("tv:", tv)

	tvOffCommand.Execute() // Выключить телевизор
	fmt.Println("tv:", tv)

	tvTurnUpVolumeCommand.Execute() // Повысить громкость
	fmt.Println("tv:", tv)

	tvTurnDownVolumeCommand.Execute() // Понизить громкость
	fmt.Println("tv:", tv)

	// Отменяем последние выполненные команды
	lightOffCommand.Undo()         // Вернуть состояние света на "включено"
	tvOffCommand.Undo()            // Вернуть состояние телевизора на "включено"
	tvTurnDownVolumeCommand.Undo() // Вернуть состояние телевизора прежняя громкость

	fmt.Println("light:", light)
	fmt.Println("tv:", tv)
}
