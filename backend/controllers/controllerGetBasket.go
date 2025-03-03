package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

func GetBasket() string {
	log.Println("Функция вызвана")
	
	data, err := ioutil.ReadFile("basket.json")
	if err != nil {
		log.Println("Файл корзины не найден, возвращаю пустой массив")
		return "[]"
	}

	var basket []map[string]interface{}
	err = json.Unmarshal(data, &basket)
	if err != nil {
		log.Println("Ошибка парсинга JSON:", err)
		return `[]`
	}
	result, _ := json.Marshal(basket)
	return string(result)
}