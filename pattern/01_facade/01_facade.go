/*
Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
https://en.wikipedia.org/wiki/Facade_pattern

Тип:		Структурный
Уровень:	Объектный

Используется для содания единого интерфейса для взаимодействия с подсистемами

Плюсы:
1. Упрощенное взаимодействие с системой
2. Сокрытие деталей реализации
3. Доступ к подсистемам

Минусы:
1.Возможное увеличение сложности при росте количеста подсистем

Примеры использования:
Бибилиотеки и фраймворки
Управление сложными системами (умный дом, и тд.)
*/
package main

// Рассмотрим пример где есть необходимость парсить obj файл и вращать объект (affine)

import "fmt"

type Facade struct {
	parse  *Parse  // подсистема 1
	affine *Affine // подсистема 2
}

func NewFacde() *Facade {
	return &Facade{parse: &Parse{}, affine: &Affine{}}
}

// подсистема 1
type Parse struct{}

func (p *Parse) ParseFile(fileName string) {
	fmt.Printf("Parse file \"%s\" is done\n", fileName)
}

// подсистема 2
type Affine struct{}

func (a *Affine) Move() {
	fmt.Println("Figure moved")
}

func (a *Affine) Transform() {
	fmt.Println("Figure transformed")
}

func (a *Affine) Scale() {
	fmt.Println("Figure scaled")
}

// main
func main() {
	facade := NewFacde()

	facade.parse.ParseFile("myFile")
	facade.affine.Move()
	facade.affine.Transform()
	facade.affine.Scale()
}
