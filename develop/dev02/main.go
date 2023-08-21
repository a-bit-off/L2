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
	"strings"
	"unicode"
)

func main() {
	str := "qwe\\45"
	res, _ := Unpack(str)
	fmt.Println(res)
}

func Unpack(str string) (string, error) {
	res := ""
	number := ""
	var slashCount int

	for i := 0; i < len(str); i++ {
		r := rune(str[i])
		if unicode.IsLetter(r) {
			if number == "" {
				res += string(r)
			} else {
				n, _ := strconv.Atoi(number)
				var size int
				if len(res) > 0 {
					size = len(res) - 1
				}
				res += strings.Repeat(res[size:], n-1)
				res += string(r)
				number = ""
			}
			slashCount = 0
		} else if unicode.IsDigit(r) {
			if slashCount == 1 {
				res += string(str[i])
			} else {
				number += string(str[i])
			}
			slashCount = 0
		} else if r == '\\' {
			slashCount++
			if slashCount == 2 {
				res += string(str[i])
				slashCount = 0
			}
		}
	}
	if number != "" {
		n, _ := strconv.Atoi(number)
		var size int
		if len(res) > 0 {
			size = len(res) - 1
		}
		res += strings.Repeat(res[size:], n-1)
	}

	return res, nil
}
