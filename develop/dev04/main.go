/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/
package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	words := []string{"Пятак", "пяткА", "тяпка", "тяпка", "листок", "слиток", "столик"}
	anagrams := findAnagrams(words)

	for k, v := range anagrams {
		fmt.Printf("%s: %v\n", k, v)
	}
}

func findAnagrams(words []string) map[string][]string {
	countWords := len(words)
	anagrams := make(map[string][]string, countWords)

	for i := 0; i < countWords; i++ {
		// переводим число в нижний регистр и сортируем (Пятка -> акптя)
		lowerWord := strings.ToLower(words[i])
		sorted := sortString(lowerWord)

		// добавляем слово в множество
		anagrams[sorted] = append(anagrams[sorted], lowerWord)
	}

	for k, v := range anagrams {
		if len(v) <= 1 { // удаляем множества из одного элемента
			delete(anagrams, k)
		} else {
			sort.Strings(v) // сортируем элементы по возрастанию
			getUnique(&v)   // удаляем дубликаты
			anagrams[k] = v
		}
	}
	return anagrams
}

func sortString(str string) string {
	strSlice := strings.Split(str, "")
	sort.Strings(strSlice)
	return strings.Join(strSlice, "")
}

func getUnique(words *[]string) {
	size := len(*words)
	unique := make([]string, 0, size)
	unique = append(unique, (*words)[0])
	for i := 1; i < size; i++ {
		if (*words)[i-1] != (*words)[i] {
			unique = append(unique, (*words)[i])
		}
	}
	*words = unique
}
