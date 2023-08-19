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

func Unpack1(str string) (string, error) {
	runes := []rune(str)
	size := len(runes)
	res := make([]rune, 0, size)
	var symbol rune
	var count []rune

	for i := 0; i < size; i++ {
		if unicode.IsDigit(runes[i]) {
			count = append(count, runes[i])
		} else {
			if len(count) == 0 {
				symbol = runes[i]
				res = append(res, runes[i])
			} else {
				c, _ := strconv.Atoi(string(count))
				repeat := strings.Repeat(string(symbol), c-1)
				res = append(res, []rune(repeat)...)

				symbol = runes[i]
				res = append(res, runes[i])

				count = []rune("")
			}
		}

	}

	return string(res), nil
}

func Unpack(str string) {
	symbols := ""
	slashes := ""
	numbers := ""
	for _, r := range str {
		fmt.Println(r)
		if unicode.IsLetter(r) {
			symbols += string(r)
		}
		//if !unicode.IsDigit(r) && r != '\\' {
		//	symbols += string(r)
		//} else if r == '\\' {
		//	slashes += string(r)
		//} else {
		//	numbers += string(r)
		//}
	}
	fmt.Println("slashes:", slashes)
	fmt.Println("symbols:", symbols)
	fmt.Println("numbers:", numbers)
}

func main() {
	str := "abcdef\\12"
	Unpack(str)
}
