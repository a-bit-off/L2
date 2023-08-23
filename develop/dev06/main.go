/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Args ...
type Args struct {
	flags Flags
	files []string
}

// Flags ...
type Flags struct {
	fields    int
	delimiter string
	separated bool
}

func main() {
	args := getArgs()

	res, err := processFile(args)
	if err != nil {
		log.Fatalln(err)
	}

	for _, v := range res {
		printLines(v)
	}

}

func getArgs() *Args {
	fieldsPtr := flag.Int("f", 0, "выбрать поля (колонки)")
	delimiterPtr := flag.String("d", "\t", "использовать другой разделитель")
	separatedPtr := flag.Bool("s", false, "только строки с разделителем")
	flag.Parse()

	files := flag.Args()[0:]

	return &Args{
		flags: Flags{
			fields:    *fieldsPtr,
			delimiter: *delimiterPtr,
			separated: *separatedPtr,
		},
		files: files,
	}
}

func processFile(args *Args) ([][]string, error) {
	var res [][]string
	for _, file := range args.files {
		lines, err := scanFile(file)
		if err != nil {
			return nil, err
		}
		filteredLines := cat(args, lines)
		res = append(res, filteredLines)
	}

	return res, nil
}

func scanFile(file string) ([]string, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, nil
}

func printLines(filteredLines []string) {
	for _, line := range filteredLines {
		fmt.Println(line)
	}
}

func cat(args *Args, lines []string) []string {
	var res []string
	sep := 0
	for _, line := range lines {

		// flag -s
		if flagSeparated(args, line, &sep) {
			continue
		}

		// flag -f
		if args.flags.fields > 0 && line != "" {
			if l := flagFields(args, line); l != "" {
				res = append(res, l)
			} else {
				sep++
			}
			continue
		}

		// flag -d
		if args.flags.delimiter != "" {
			res = append(res, flagDelimiter(args, line))
			continue
		}

		res = append(res, line)
	}
	return res
}

func flagSeparated(args *Args, line string, sep *int) bool {
	if args.flags.separated {
		if line == "" {
			*sep++
			if *sep > 1 {
				return true
			}
		} else {
			*sep = 0
		}
	}
	return false
}

func flagFields(args *Args, line string) string {
	f := args.flags.fields - 1

	split := strings.Split(line, " ")
	if f >= len(split) {
		return ""
	}

	return split[f]
}

func flagDelimiter(args *Args, line string) string {
	newStr, _ := strconv.Unquote(`"` + strings.ReplaceAll(line, " ", args.flags.delimiter) + `"`)
	return newStr
}
