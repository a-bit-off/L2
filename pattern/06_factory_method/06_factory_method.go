/*
Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
https://en.wikipedia.org/wiki/Factory_method_pattern

Тип:		Порождающий
Уровень:	Класса

Паттерн Factory Method полезен, когда система должна оставаться легко расширяемой путем добавления объектов новых типов.
Этот паттерн является основой для всех порождающих паттернов и может легко трансформироваться под нужды системы.
По этому, если перед разработчиком стоят не четкие требования для продукта или не ясен способ организации взаимодействия
между продуктами, то для начала можно воспользоваться паттерном Factory Method, пока полностью не сформируются все
требования.

Паттерн Factory Method применяется для создания объектов с определенным интерфейсом, реализации которого предоставляются
потомками. Другими словами, есть базовый абстрактный класс фабрики, который говорит, что каждая его наследующая фабрика
олжна реализовать такой-то метод для создания своих продуктов.

Пример: К нам приходят файлы трех расширений .txt, .png, .doc. В зависимости от расширения файла мы должны сохранять
его в одном из каталогов /file/txt/, /file/png/ и /file/doc/. Значит, у нас будет файловая фабрика с параметризированным
фабричным методом, принимающим путь к файлу, который нам нужно сохранить в одном из каталогов. Этот фабричный метод
озвращает нам объект, используя который мы можем манипулировать с нашим файлом (сохранить, посмотреть тип и каталог
для сохранения). Заметьте, мы никак не указываем какой экземпляр объекта-продукта нам нужно получить, это делает
фабричный метод путем определения расширения файла и на его основе выбора подходящего класса продукта. Тем самым,
если наша система будет расширяться и доступных расширений файлов станет, например 25, то нам всего лишь нужно
будет изменить фабричный метод и реализовать классы продуктов.

Требуется для реализации:

 1. Базовый абстрактный класс Creator, описывающий интерфейс, который должна реализовать конкретная фабрика для
    производства продуктов. Этот базовый класс описывает фабричный метод.
 2. Базовый класс Product, описывающий интерфейс продукта, который возвращает фабрика. Все продукты возвращаемые
    фабрикой должны придерживаться единого интерфейса.
 3. Класс конкретной фабрики по производству продуктов ConcreteCreator. Этот класс должен реализовать фабричный метод;
 4. Класс реального продукта ConcreteProductA;
 5. Класс реального продукта ConcreteProductB;
 6. Класс реального продукта ConcreteProductC.

Factory Method отличается от Abstract Factory, тем, что Abstract Factory производит семейство объектов, эти объекты
разные, обладают разными интерфейсами, но взаимодействуют между собой. В то время как Factory Method производит
продукты придерживающиеся одного интерфейса и эти продукты не связаны между собой, не вступают во взаимодействие.
*/

package main

import "fmt"

// Предположим, у нас есть приложение для обработки документов, где каждый документ может иметь свой собственный формат.
// Мы хотим иметь возможность создавать различные типы документов, такие как PDF, Word и Excel.

// Базовый класс Product, описывающий интерфейс продукта, который возвращает фабрика. Все продукты возвращаемые
// фабрикой должны придерживаться единого интерфейса.
type Document interface {
	Open() Document
	Save() Document
}

// PDF
type PDFDocument struct{}

func (d *PDFDocument) Open() Document {
	fmt.Println("PDF open!")
	return d
}

func (d *PDFDocument) Save() Document {
	fmt.Println("PDF save!")
	return d
}

// Word
type WordDocument struct{}

func (d *WordDocument) Open() Document {
	fmt.Println("Word open!")
	return d
}

func (d *WordDocument) Save() Document {
	fmt.Println("Word save!")
	return d
}

// Excel
type ExcelDocument struct{}

func (d *ExcelDocument) Open() Document {
	fmt.Println("Excel open!")
	return d
}

func (d *ExcelDocument) Save() Document {
	fmt.Println("Excel save!")
	return d
}

// Базовый абстрактный класс Creator, описывающий интерфейс, который должна реализовать конкретная фабрика для
// производства продуктов. Этот базовый класс описывает фабричный метод.
type DocumentFactory interface {
	CreateDocument() Document
}

// Класс конкретной фабрики по производству продуктов ConcreteCreator. Этот класс должен реализовать фабричный метод;

// PDF FACTORY
type PDFFactory struct{}

func (f *PDFFactory) CreateDocument() Document {
	return &PDFDocument{}
}

// Word FACTORY
type WordFactory struct{}

func (f *WordFactory) CreateDocument() Document {
	return &WordDocument{}
}

// Excel FACTORY
type ExcelFactory struct{}

func (f *ExcelFactory) CreateDocument() Document {
	return &ExcelDocument{}
}

func main() {
	pdfFactory := &PDFFactory{}
	pdfDocument := pdfFactory.CreateDocument()
	pdfDocument.Open()
	pdfDocument.Save()

	wordFactory := &WordFactory{}
	wordDocument := wordFactory.CreateDocument()
	wordDocument.Open().Save()

	excelFactory := &ExcelFactory{}
	excelDocument := excelFactory.CreateDocument().Open()
	excelDocument.Save()

}
