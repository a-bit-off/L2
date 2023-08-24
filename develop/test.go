package main

import (
	"fmt"
	"time"
)

func main() {
	dateString := "01-02-2022" // Пример строки с датой в формате "день-месяц-год"

	// Преобразование строки в дату
	date, err := time.Parse("02-01-2006", dateString)
	if err != nil {
		fmt.Println("Ошибка преобразования строки в дату:", err)
		return
	}

	// Форматирование даты в нужный формат "день месяц год"
	formattedDate := date.Format("02 January 2006")
	fmt.Println("Преобразованная дата:", formattedDate)
}