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

// Args ...
type Args struct {
	lcK int
	lcN bool
	lcR bool
	lcU bool

	ucM bool
	lcB bool
	lcC bool
	lcH bool

	filePath string
}

func main() {
	args, err := getArgs()
	if err != nil {
		log.Fatalln(err)
	}
	data, err := getData(args.filePath)
	if err != nil {
		log.Fatalln(err)
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

	M := flag.Bool("M", false, "сортировать по названию месяца")
	b := flag.Bool("b", false, "игнорировать хвостовые пробелы")
	c := flag.Bool("c", false, "проверять отсортированы ли данные")
	h := flag.Bool("h", false, "сортировать по числовому значению с учётом суффиксов")
	flag.Parse()

	if len(os.Args) == 1 {
		return nil, fmt.Errorf("укажите файл")
	}
	filePath := os.Args[len(os.Args)-1]

	return &Args{
		lcK: *k,
		lcN: *n,
		lcR: *r,
		lcU: *u,

		ucM: *M,
		lcB: *b,
		lcC: *c,
		lcH: *h,

		filePath: filePath,
	}, nil
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

	if args.ucM {
		res = flagUcM(data)
	}
	if args.lcB {
		res = flagLcB(data)
	}
	if args.lcC {
		c := flagLcC(data)
		fmt.Printf("sort: %s: flag -c: %t\n", args.filePath, c)
		if !c {
			return nil
		}

	}
	if args.lcH {
		res = flagLcH(data)
	}

	return res
}

// main flags
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

	for k := range myMap {
		res = append(res, k)
	}
	return res
}

// bonus flags
func flagUcM(data []string) []string {
	sort.Slice(data, func(i, j int) bool {
		return parseMonth(data[i]) < parseMonth(data[j])
	})
	return data
}

func parseMonth(s string) int {
	var month string
	fmt.Sscanf(s, "%s", &month)

	switch month {
	case "Январь":
		return 1
	case "Февраль":
		return 2
	case "Март":
		return 3
	case "Апрель":
		return 4
	case "Май":
		return 5
	case "Июнь":
		return 6
	case "Июль":
		return 7
	case "Август":
		return 8
	case "Сентябрь":
		return 9
	case "Октябрь":
		return 10
	case "Ноябрь":
		return 11
	case "Декабрь":
		return 12
	}
	return 0
}

func flagLcB(data []string) []string {
	sort.Slice(data, func(i, j int) bool {
		return strings.TrimSpace(data[i]) < strings.TrimSpace(data[j])
	})
	return data
}

func flagLcC(data []string) bool {
	return sort.SliceIsSorted(data, func(i, j int) bool {
		return data[i] < data[j]
	})
}

func flagLcH(data []string) []string {
	sort.Slice(data, func(i, j int) bool {
		return parseSize(data[i]) < parseSize(data[j])
	})
	return data
}

func parseSize(s string) float64 {
	var size float64
	var suffix string
	fmt.Sscanf(s, "%f%s", &size, &suffix)
	switch suffix {
	case "K":
		size *= 1e3
	case "M":
		size *= 1e6
	case "G":
		size *= 1e9
	case "T":
		size *= 1e12
	}
	return size
}
