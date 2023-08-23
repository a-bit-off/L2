package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

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
"io"
"os"
"strings"
)

func grep(pattern string, lines []string, flags *Flags) []string {
	var filteredLines []string

	for i, line := range lines {
		match := strings.Contains(line, pattern)

		if flags.ignoreCase {
			match = strings.Contains(strings.ToLower(line), strings.ToLower(pattern))
		}

		if (flags.invert && !match) || (!flags.invert && match) {
			if flags.count {
				flags.matchCount++
			} else {
				if flags.lineNum {
					line = fmt.Sprintf("%d:%s", i+1, line)
				}
				filteredLines = append(filteredLines, line)
			}
		}
	}

	return filteredLines
}

type Flags struct {
	after      int
	before     int
	context    int
	count      bool
	ignoreCase bool
	invert     bool
	fixed      bool
	lineNum    bool
	matchCount int
}

func main() {
	afterPtr := flag.Int("A", 0, "Print +N lines after match")
	beforePtr := flag.Int("B", 0, "Print +N lines before match")
	contextPtr := flag.Int("C", 0, "Print ±N lines around match")
	countPtr := flag.Bool("c", false, "Print count of matching lines")
	ignoreCasePtr := flag.Bool("i", false, "Ignore case")
	invertPtr := flag.Bool("v", false, "Invert match")
	fixedPtr := flag.Bool("F", false, "Fixed string match")
	lineNumPtr := flag.Bool("n", false, "Print line numbers")

	flag.Parse()

	flags := Flags{
		after:      *afterPtr,
		before:     *beforePtr,
		context:    *contextPtr,
		count:      *countPtr,
		ignoreCase: *ignoreCasePtr,
		invert:     *invertPtr,
		fixed:      *fixedPtr,
		lineNum:    *lineNumPtr,
	}

	pattern := flag.Arg(0)
	files := flag.Args()[1:]

	if pattern == "" {
		fmt.Println("Pattern is required")
		return
	}

	if len(files) == 0 {
		filteredLines := grep(pattern, readLines(os.Stdin), &flags)
		printLines(filteredLines, flags)
	} else {
		for _, file := range files {
			f, err := os.Open(file)
			if err != nil {
				fmt.Printf("Failed to open file: %s\n", file)
				continue
			}

			filteredLines := grep(pattern, readLines(f), &flags)
			printLines(filteredLines, flags)

			f.Close()
		}
	}

	if flags.count {
		fmt.Printf("Match count: %d\n", flags.matchCount)
	}
}

func readLines(r io.Reader) []string {
	var lines []string
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func printLines(lines []string, flags Flags) {
	for _, line := range lines {
		fmt.Println(line)
	}
}