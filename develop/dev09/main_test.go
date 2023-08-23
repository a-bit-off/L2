package main

import (
	"fmt"
	"os"
	"testing"
)

func TestParseHTML_Positive(t *testing.T) {
	// Подготовка тестовых данных
	url := "https://example.com"
	i := 1

	// Выполнение функции
	parseHTML(url, i)

	// Проверка результатов
	fileName := fmt.Sprintf("copied_page%d.html", i)
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		t.Errorf("Файл %s не был создан", fileName)
	}
}

func TestParseHTML_Negative(t *testing.T) {
	// Подготовка тестовых данных
	url := "https://nonexistenturl.com"
	i := 2

	// Выполнение функции
	err := parseHTML(url, i)
	fileName := fmt.Sprintf("copied_page%d.html", i)
	if err == nil {
		t.Errorf("Файл %s был некорректно создан", fileName)
	}

}
