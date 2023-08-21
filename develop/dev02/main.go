/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func main() {
	str := "45"
	res, _ := unpackString(str)
	fmt.Println(res)
}

func unpackString(str string) (string, error) {
	buf := make([]byte, 0, len(str))
	var number []byte
	var slashCount int

	for _, r := range str {
		if unicode.IsLetter(r) {
			if len(number) == 0 {
				buf = append(buf, byte(r)) // добавляем в результирующую строку символ, если не нужно дублировать
			} else {
				if err := addRepeat(&buf, &number); err != nil { // дублируем последний символ
					return "", err
				}
				buf = append(buf, byte(r)) // добавляем текущий сивол
			}
			slashCount = 0
		} else if unicode.IsDigit(r) {
			if slashCount == 1 { // если числу предшествует слеш то добавялем число в строку
				buf = append(buf, byte(r))
			} else { // если же слешей нет, то парсим число
				number = append(number, byte(r))
			}
			slashCount = 0
		} else if r == '\\' {
			slashCount++
			if slashCount == 2 { // если встречаем два слеша, то добавялем слеш в строку
				buf = append(buf, byte(r))
				slashCount = 0
			}
		}
	}

	if len(number) > 0 { // для случаев "abcd12" дублируем последний символ "d"
		if err := addRepeat(&buf, &number); err != nil {
			return "", err
		}
	}

	return string(buf), nil
}

func addRepeat(buf, number *[]byte) error {
	n, err := strconv.ParseInt(string(*number), 10, 64) // парсим число
	if err != nil {
		return err
	}
	size := len(*buf)
	if size == 0 {
		return nil
	} else if size > 0 {
		size--
	}

	for i := 0; i < int(n-1); i++ { // дублируем символ
		*buf = append(*buf, (*buf)[size])
	}

	*number = nil
	return nil
}
