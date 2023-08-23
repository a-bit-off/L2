/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	urls, err := getUrls(os.Args[1:]...)
	if err != nil {
		log.Fatalln(err)
	}

	for i, url := range urls {
		err = parseHTML(url, i)
		if err != nil {
			log.Fatalln(err)
		}
	}
}

func getUrls(urls ...string) ([]string, error) {
	var links []string
	if len(urls) < 1 {
		return nil, errors.New("go-wget: missing URL")
	}
	links = urls
	return links, nil
}

func parseHTML(url string, i int) error {
	// Выполняем GET-запрос
	response, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("ошибка при выполнении GET-запроса: %s", err)
	}
	defer response.Body.Close()

	// Читаем содержимое страницы
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return fmt.Errorf("ошибка при чтении содержимого: %s", err)
	}

	// Записываем содержимое страницы в файл
	fileName := fmt.Sprintf("copied_page%d.html", i)
	err = ioutil.WriteFile(fileName, body, os.ModePerm)
	if err != nil {
		return fmt.Errorf("ошибка при записи файла: %s", err)
	}

	return nil
}
