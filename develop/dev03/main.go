/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

- k — указание колонки для сортировки
- n — сортировать по числовому значению
- r — сортировать в обратном порядке
- u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

// lowercase lc
// uppercase uc

type Args struct {
	lcK      int
	lcN      bool
	lcR      bool
	lcU      bool
	filePath string
}

func main() {
	args, err := getArgs()
	if err != nil {
		log.Fatalln(err)
	}
	data, err := getData(args.filePath)
	if err != nil {
		log.Println(err)
	}

	res := mySort(args, data)
	for _, r := range res {
		fmt.Println(r)
	}
}

func getArgs() (*Args, error) {
	k := flag.Int("k", 0, "указание колонки для сортировки")
	n := flag.Bool("n", false, "сортировать по числовому значению")
	r := flag.Bool("r", false, "сортировать в обратном порядке")
	u := flag.Bool("u", false, "не выводить повторяющиеся строки")
	flag.Parse()

	if len(os.Args) == 1 {
		return nil, fmt.Errorf("Укажите файл!")
	}
	filePath := os.Args[len(os.Args)-1]
	//filePath := "files/case1_k.txt"

	return &Args{lcK: *k, lcN: *n, lcR: *r, lcU: *u, filePath: filePath}, nil
}

func getData(filePath string) ([]string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	data := make([]string, 0, 10)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	if scanner.Err() != nil {
		return nil, err
	}
	return data, nil
}

func mySort(args *Args, data []string) []string {
	res := data

	if args.lcK > 0 {
		res = flagLcK(res, args)
	} else if args.lcR {
		res = flagLcR(res)
	} else if args.lcN {
		res = flagLcN(res)
	}

	if args.lcU {
		res = flagLcU(data)
	}
	return res
}

func flagLcK(data []string, args *Args) []string {
	position := args.lcK
	if position <= 0 {
		return nil
	}
	position--
	sort.Slice(data, func(i, j int) bool {
		di := strings.Split(data[i], " ")
		if len(di) <= position {
			if args.lcR {
				return false
			}
			return true
		}

		dj := strings.Split(data[j], " ")
		if len(dj) <= position {
			if args.lcR {
				return true
			}
			return false
		}
		if args.lcR {
			return di[position] > dj[position]
		}
		return di[position] < dj[position]
	})
	return data
}

func flagLcN(data []string) []string {
	sort.Slice(data, func(i, j int) bool {
		return data[i] < data[j]
	})
	return data
}

func flagLcR(data []string) []string {
	sort.Slice(data, func(i, j int) bool {
		return data[i] > data[j]
	})
	return data
}

func flagLcU(data []string) []string {
	size := len(data)
	res := make([]string, 0, size)
	myMap := make(map[string]int, size)

	for i := 0; i < size; i++ {
		myMap[data[i]] = i
	}

	for k, _ := range myMap {
		res = append(res, k)
	}
	return res
}
