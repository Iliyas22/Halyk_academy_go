package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Структура для хранения данных о породе
type Breed struct {
	Name   string `json:"name"`
	Origin string `json:"origin"`
	// Добавьте другие поля по мере необходимости
}

func main() {
	url := "https://catfact.ninja/breeds" // Замените на URL вашего API

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Ошибка при отправке запроса:", err)
		return
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Ошибка при чтении ответа:", err)
		return
	}

	var breeds []Breed
	err = json.Unmarshal(body, &breeds)
	if err != nil {
		fmt.Println("Ошибка при разборе JSON:", err)
		return
	}

	// Создайте карту для группировки по стране происхождения
	breedsByOrigin := make(map[string][]string)

	// Заполните карту
	for _, breed := range breeds {
		breedsByOrigin[breed.Origin] = append(breedsByOrigin[breed.Origin], breed.Name)
	}

	// Выведите результаты
	for origin, breedNames := range breedsByOrigin {
		fmt.Printf("Породы из страны %s: %v\n", origin, breedNames)
	}
}
