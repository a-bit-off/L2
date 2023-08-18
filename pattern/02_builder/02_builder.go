/*
	Реализовать паттерн «строитель».

Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.

	https://en.wikipedia.org/wiki/Builder_pattern

Тип:		Порождающий
Уровень: 	Объектный

Паттерн "Строитель" (Builder) является порождающим паттерном проектирования, который позволяет создавать
сложные объекты шаг за шагом. Он позволяет отделить процесс создания объекта от его представления,
что позволяет получить разные представления одного и того же объекта.

Кока-Кола производит сложный продукт, состоящий из 4 частей (крышка, бутылка, этикетка, напиток),
которые должны быть применены в нужном порядке. Нельзя вначале взять крышку, бутылку, завинтить крышку,
а потом пытаться налить туда напиток. Для реализации объекта, бутылки Кока-Колы, которая поставляется клиенту,
нам нужен паттерн Builder.

Требуется для реализации:
 1. Базовый абстрактный класс Builder, который описывает интерфейс строителя, те команды, которые он обязан выполнять;
 2. Класс Director, который будет распоряжаться строителем и отдавать ему команды в нужном порядке,
    а строитель будет их выполнять;
 3. Класс сложного объекта Product.
 4. Класс ConcreteBuilder, который реализует интерфейс строителя и взаимодействует со сложным объектом;

Применение паттерна "Строитель" полезно в следующих случаях:
1. Когда требуется создание сложных объектов с различными вариациями и конфигурациями.
2. Когда необходимо избежать загрязнения конструктора множеством параметров.
3. Когда нужно обеспечить последовательное создание и настройку объектов.
4. Когда требуется создание объектов с различными представлениями.

Плюсы:
- Упрощение процесса создания сложных объектов.
- Разделение процесса конструирования и представления объекта.
- Позволяет создавать различные представления одного и того же объекта.

Минусы:
- Увеличение количества классов и сложности кода.
- Возможное увеличение времени разработки из-за необходимости создания дополнительных классов.

Примеры использования:
1. Построение графических интерфейсов:
В графических библиотеках, таких как Qt или Swing, паттерн "Строитель" может использоваться для создания сложных
пользовательских интерфейсов, где каждый компонент может иметь различные свойства и настройки.
2. Конструирование объектов баз данных:
В ORM (Object-Relational Mapping) системах, паттерн "Строитель" может быть использован для создания запросов к
базе данных, где каждый компонент запроса может быть добавлен поэтапно.
*/
package main

// реализуем паттерн на пирмере сборщика компютеров

import "fmt"

// Класс сложного объекта Product.
type Computer struct {
	CPU string
	RAM int
	MB  string
}

// Базовый абстрактный класс Builder, который описывает интерфейс строителя, те команды, которые он обязан выполнять;
type ComputerBuilderI interface {
	CPU(val string) ComputerBuilderI
	RAM(val int) ComputerBuilderI
	MB(val string) ComputerBuilderI

	Build() Computer
}

// Класс ConcreteBuilder, который реализует интерфейс строителя и взаимодействует со сложным объектом;
type ComputerBuilder struct {
	cpu string
	ram int
	mb  string
}

func NewComputerBuilder() ComputerBuilderI {
	return ComputerBuilder{}
}

func (cb ComputerBuilder) CPU(val string) ComputerBuilderI {
	cb.cpu = val
	return cb
}

func (cb ComputerBuilder) RAM(val int) ComputerBuilderI {
	cb.ram = val
	return cb
}

func (cb ComputerBuilder) MB(val string) ComputerBuilderI {
	cb.mb = val
	return cb
}

func (cb ComputerBuilder) Build() Computer {
	return Computer{
		CPU: cb.cpu,
		RAM: cb.ram,
		MB:  cb.mb,
	}
}

// Второй класс, который реализует интерфейс
type OfficeComputerBuilder struct {
	ComputerBuilder
}

func NewOfficeComputerBuilder() ComputerBuilderI {
	return ComputerBuilder{}.CPU("intel pentium III").RAM(2).MB("kingston")
}

func (b *OfficeComputerBuilder) Build() Computer {
	return Computer{
		CPU: b.cpu,
		RAM: b.ram,
		MB:  b.mb,
	}
}

// Класс Director, который будет распоряжаться строителем и отдавать ему команды в нужном порядке,
// а строитель будет их выполнять;
type Director struct {
	computerBuilderI ComputerBuilderI
}

func NewDirector(b ComputerBuilderI) *Director {
	return &Director{
		computerBuilderI: b,
	}
}

func (d *Director) SetComputer(build ComputerBuilderI) {
	d.computerBuilderI = build
}

func (d *Director) BuildComputer() Computer {
	return d.computerBuilderI.Build()
}

// Main
func main() {
	cb := NewComputerBuilder()
	computer := cb.CPU("core i3").RAM(8).MB("gigabyte").Build()
	fmt.Println("custom pc:", computer)

	officePC := NewOfficeComputerBuilder().Build()
	fmt.Println("default office pc:", officePC)

	myOfficePC := NewOfficeComputerBuilder().RAM(8).Build()
	fmt.Println("custom office pc:", myOfficePC)

	director := NewDirector(NewOfficeComputerBuilder())
	com := director.BuildComputer()
	fmt.Println("default office pc made by director:", com)

	director.SetComputer(NewComputerBuilder().CPU("core i9").RAM(32).MB("gigabyte"))
	com = director.BuildComputer()
	fmt.Println("custom pc made by director:", com)
}
